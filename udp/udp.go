package udp
/*
  Handles listening for incoming gossip requests, group joins, and quits
  Adapted from 
    https://github.com/bootic/bootic_data_collector/blob/master/udp/udp.go
*/

import (
  "net"
  "log"
)

type Daemon struct {
  Conn *net.UDPConn
}

func NewDaemon(hostPort string) (daemon *Daemon, err error) {
  conn, err := createUDPListener(hostPort)

  if err != nil {
    return
  }

  daemon = &Daemon {
    Conn: conn,
  }

  log.Println("Daemon created!")
  return
}

func (self *Daemon) ReceiveDatagrams() {
  for {
    buffer := make([]byte, 1024)

    // Fancy error checking one-liner
    if c, addr, err := self.Conn.ReadFromUDP(buffer); err != nil {
      log.Printf("%d byte datagram from %s with error %s\n", c, addr.String(), err.Error())
      return
    } else {
      log.Println("Data received: ", string(buffer[:c]))
    }
  }
}

// Send message to address
func (self *Daemon) Gossip(message, address string) (err error){
  var udpaddr *net.UDPAddr
  if udpaddr, err = net.ResolveUDPAddr("udp4", address); err != nil {
    return
  }

  var conn *net.UDPConn
  conn, err = net.DialUDP("udp4", nil, udpaddr)
  log.Printf("Sending '%s' to %s..", message, udpaddr)
  if _, err = conn.Write([]byte(message)); err != nil {
    log.Panic("Writing to UDP", err)
  }

  return
}


func createUDPListener(hostPort string) (conn *net.UDPConn, err error) {

  var udpaddr *net.UDPAddr
  if udpaddr, err = net.ResolveUDPAddr("udp4", hostPort); err != nil {
    return
  }

  conn, err = net.ListenUDP("udp4", udpaddr)

  return
}
