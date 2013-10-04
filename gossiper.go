package main

import (
)

// Maintain dictionary of machines
// Each entry counts number of "heartbeats" since we heard from that machine
// If the number of heartbeats crosses a threshold, we know it is unresponsive

func main() {
  // Generate an ID to join the group with
  // Join the group, receive current group list

  // LOOP for every heartbeat:
    // Choose a random machine to gossip to
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




