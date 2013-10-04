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

  // Get command-line arguments for setting up group options
  var address = flag.String("a", "",
    "Address of a machine in the group to join. Blank if this is a new group.")
  flag.Parse()
  log.Println(*address)


  daemon, err := udp.NewDaemon("localhost:4567")
  if err != nil {
    log.Panic("Daemon creation", err)
  }

  log.Println(daemon)

  /* 
    use this to test udp:
    echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon.ReceiveDatagrams()


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




