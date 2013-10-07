package main

import (
  "./data"
  "./udp"
  "flag"
  "log"
  "time"
  "math/rand"
)


// Maintain dictionary of machines
// Each entry counts number of "heartbeats" since we heard from that machine
// If the number of heartbeats crosses a threshold, we know it is unresponsive

func main() {

  // Parse command-line arguments
  var (
    udpHost       string
    groupMember   string
  )

  flag.StringVar(&udpHost, "udphost", "localhost:4567", "host:port to bind for UDP listener")
  flag.StringVar(&groupMember, "g", "", "address of an existing group member")
  flag.Parse()

  log.Println("Start server on:", udpHost)

  groupList := make(map[string]*data.GroupMember)

  // Determine the heartbeat duration time
  heartbeatInterval := 100 * time.Millisecond

  /* 
    use this to test (from command line for now):
      echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon, err := udp.NewDaemon(udpHost)
  if err != nil {
    log.Panic("Daemon creation", err)
  }

  // Join the group by broadcasting a dummy message
  // TODO what if the packet got dropped? rebroadcast after a timeout
  if groupMember != "" {
    udp.SendMessage("JOIN", groupMember)
    log.Println("GOSSIP","Gossiping new member to the group") // TODO use mp1 logger
  }

  go daemon.ReceiveDatagrams()

  //TODO: Figure out how to move list to Daemon and use the ticker and listen to UDP in parallel
  for {
    //Get random member , increment current members
    heartbeatAndGossip(groupList)
    log.Println("TICK", time.Now())
    time.Sleep(heartbeatInterval)
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

// Send a packet about subject to receiver over UDP. If subject is nil, send a dummy message
func gossip(subject, receiver *data.GroupMember) {
  // The message we are sending over UDP
  msg := "GOSSIP"
  defer udp.SendMessage(msg, receiver.Address)

  // Determine if there is a subject or not, convert to string if so
  if subject == nil {
    return
  }
  msg = data.Marshal(subject)
}


// Increment all heartbeat values in the member list
// also send a random member A's information to a random member B
func heartbeatAndGossip(groupList map[string]*data.GroupMember) {
  // Nobody in the list yet
  if len(groupList) < 1 {
    return
  }

  // There is only one other machine, send them a dummy message
  if len(groupList) == 1 {
    // is this the only way to get the first element of a map?
    for _, receiver := range groupList {
      gossip(nil, receiver)
    }
    return
  }

  subjectIndex, receiverIndex := distinctRandoms(len(groupList))
  var subject, receiver *data.GroupMember
  i := 0
  for _, currMember := range groupList {
    if subjectIndex == i { subject = currMember }
    if receiverIndex == i { receiver = currMember }
    currMember.IncrementHeartBeat()
    i++
  }
  gossip(subject, receiver)
}

