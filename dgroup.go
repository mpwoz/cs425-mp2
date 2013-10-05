package main

import (
  "flag"
  "log"
  "./udp"
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

  /* 
    use this to test (from command line for now):
      echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon, err := udp.NewDaemon(udpHost)
  if err != nil {
    log.Panic("Daemon creation", err)
  }
  if groupMember != "" {
    daemon.Gossip("hello, from daemon", groupMember)
  }
  //daemon.ReceiveDatagrams()
  for {
  }
  log.Println("Done")


  // Generate an ID to join the group with (may not be possible)
  // Join the group, receive current group list (with self included)

  // LOOP for every heartbeat:
    // Choose a random machine from list to gossip to
    // Broadcast current group state to chosen machine
    // Increment heartbeat counters for all machines (zero own counter)
}

// Handle join request from a new machine
// Add it to the list
// Send it the list

// Handle quit request from any machine
// Set heartbeats for that machine to -1 (code for quitting)

// Handle incoming gossip, g, from any machine
// For each item in g:
  // Add it if it doesn't exist
  // if heartbeats are lower than in our list:
  // Update our list with the lower (more current) heartbeats




