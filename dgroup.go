package main

import (
  "./udp"
  "flag"
  "log"
  "time"
)


// Maintain dictionary of machines
// Each entry counts number of "heartbeats" since we heard from that machine
// If the number of heartbeats crosses a threshold, we know it is unresponsive

func main() {

  // Parse command-line arguments
  var (
    listenPort       string
    groupMember   string
  )

  flag.StringVar(&listenPort, "listen", "4567", "port to bind for UDP listener")
  flag.StringVar(&groupMember, "g", "", "address of an existing group member")
  flag.Parse()

  log.Println("Start server on port", listenPort)

  // Determine the heartbeat duration time
  heartbeatInterval := 100 * time.Millisecond

  /* 
    use this to test (from command line for now):
      echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon, err := udp.NewDaemon(listenPort)
  if err != nil {
    log.Panic("Daemon creation", err)
  }

  // Join the group by broadcasting a dummy message
  // TODO what if the packet got dropped? rebroadcast after a timeout
  firstInGroup := groupMember == ""
  if !firstInGroup {
    daemon.JoinGroup(groupMember)
    log.Println("JOIN","Gossiping new member to the group") // TODO use mp1 logger
  }

  go daemon.ReceiveDatagrams(firstInGroup)

  for {
    //Get random member , increment current members
    daemon.HeartbeatAndGossip()
    time.Sleep(heartbeatInterval)
  }
}


