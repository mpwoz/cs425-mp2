package udp
/*
  Handles listening for incoming gossip requests, group joins, and quits
  Adapted from 
    https://github.com/bootic/bootic_data_collector/blob/master/udp/udp.go
*/

import (
  "../data"
  "log"
  "math/rand"
  "net"
  "strings"
  "time"
)
const (
  heartbeatThreshold = 50
)


type Daemon struct {
  Conn *net.UDPConn
  Port string
  MemberList map[string]*data.GroupMember
}

func NewDaemon(port string) (daemon *Daemon, err error) {
  hostPort := net.JoinHostPort("localhost", port)
  log.Printf("Creating daemon at %s\n", hostPort)
  conn, err := createUDPListener(hostPort)

  if err != nil {
    return
  }

  daemon = &Daemon {
    Conn: conn,
    Port: port,
    MemberList: make(map[string]*data.GroupMember),
  }

  log.Println("Daemon created!")
  return
}

func (self *Daemon) ReceiveDatagrams(joinGroupOnConnection bool) {
  for {
    buffer := make([]byte, 1024)
    c, addr, err := self.Conn.ReadFromUDP(buffer)
    if err != nil {
      log.Printf("%d byte datagram from %s with error %s\n", c, addr.String(), err.Error())
      return
    }

    portmsg := strings.SplitN(string(buffer[:c]), "<PORT>", 2)
    port, msg := portmsg[0], portmsg[1]
    senderAddr := net.JoinHostPort(addr.IP.String(), port)

    //log.Printf("Data received from %s: %s", senderAddr, msg)

    self.handleMessage(msg, senderAddr, &joinGroupOnConnection)
  }
}


func (self *Daemon) handleMessage(msg, sender string, joinSenderGroup *bool) {
  fields := strings.SplitN(msg, "|%|", 2)
  switch fields[0] {
    case "JOIN":
      self.addNewMember(sender)
      if *joinSenderGroup {
        *joinSenderGroup = false
        self.JoinGroup(sender)
      }
    case "GOSSIP":
      self.handleGossip(sender, fields[1])
  }
}

func (self *Daemon) handleGossip(senderAddr, subject string) {
  // Reset the counter for the sender
  // TODO add sender if it doesn't exist yet
  for id,member := range self.MemberList {
    if strings.Contains(id, senderAddr) {
      //log.Printf("Reset %s, %s\n", id, senderAddr)
      member.SetHeartBeat(0)
    }
  }

  // Update the counter for the subject
  subjectMember := data.Unmarshal(subject)
  if subjectMember == nil {
    return
  }

  curr := self.MemberList[subjectMember.Id]
  if curr == nil {
    self.MemberList[subjectMember.Id] = subjectMember
  } else {
    if curr.Heartbeat > subjectMember.Heartbeat {
      curr.SetHeartBeat(subjectMember.Heartbeat)
    }
  }

}

func (self *Daemon) addNewMember(address string) (newMember *data.GroupMember){
  now := time.Now().UTC()
  machineId := address + "###" + now.String()
  newMember = data.NewGroupMember(machineId, address, 0)
  log.Printf("Created new member with IP: %s", address)
  self.MemberList[machineId] = newMember
  return
}

func (self *Daemon) JoinGroup(address string) (err error) {
  return self.sendMessageWithPort("JOIN", address)
}

func (self *Daemon) Gossip(subject,receiver *data.GroupMember) (err error) {
  // The message we are sending over UDP, subject can be nil
  msg := "GOSSIP|%|" + data.Marshal(subject)
  return self.sendMessageWithPort(msg, receiver.Address)
}

func (self *Daemon) sendMessageWithPort(msg, address string) (err error) {
  msg = self.Port + "<PORT>" + msg
  return sendMessage(msg, address)
}


// Increment all heartbeat values, and gossip a random machine A to another random B
func (self *Daemon) HeartbeatAndGossip() {
  // Nobody in the list yet
  if len(self.MemberList) < 1 {
    return
  }

  receiverIndex := rand.Int() % len(self.MemberList)
  var receiver *data.GroupMember
  i := 0
  for _, currMember := range self.MemberList {
    if receiverIndex == i { receiver = currMember }
    currMember.IncrementHeartBeat()
    if currMember.Heartbeat > heartbeatThreshold {
      log.Println("MACHINE DEAD!", currMember.Id)
    }
    i++
  }
  for _, subject := range self.MemberList {
    if subject.Id == receiver.Id {
      self.Gossip(nil, receiver)
    }
    self.Gossip(subject, receiver)
  }
}


// Get two random numbers that aren't the same
// TODO there's probably a better (quicker) way to do this 
func distinctRandoms(max int) (a, b int) {
  a = rand.Int() % max
  for b = a; b == a; {
    b = rand.Int() % max
  }
  return
}


func sendMessage(message, address string) (err error) {
  var raddr *net.UDPAddr
  if raddr, err = net.ResolveUDPAddr("udp", address); err != nil {
    log.Panic(err)
  }

  var con *net.UDPConn
  con, err = net.DialUDP("udp", nil, raddr)
  //log.Printf("Sending '%s' to %s..", message, raddr)
  if _, err = con.Write([]byte(message)); err != nil {
    log.Panic("Writing to UDP:", err)
  }

  return
}


func createUDPListener(hostPort string) (conn *net.UDPConn, err error) {

  var udpaddr *net.UDPAddr
  if udpaddr, err = net.ResolveUDPAddr("udp", hostPort); err != nil {
    return
  }

  conn, err = net.ListenUDP("udp", udpaddr)

  return
}
