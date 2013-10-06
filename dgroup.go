package main

import (
  "flag"
  "log"
  "time"
  "./udp"
  "./data/GroupMember"
  "os"
  "./logger"
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

  //Use the actual hostname instead of localhost - ensure uniqueness when transmitting groupmemberdata
  localhost, err := os.Hostname()
  if err != nil {
    fmt.Printf("Oops: %v\n", err)
    return
}
  flag.StringVar(&udpHost, "udphost", localhost+":4567", "host:port to bind for UDP listener")
  flag.StringVar(&groupMember, "g", "", "address of an existing group member")
  flag.Parse()

  log.Println("Start server on:", udpHost)

  //Maintain a dictionary of machines
  var groupList Map[String]GroupMember
  
  // Determine the heartbeat duration time
  duration int
  duration = 5
  
  //Unique Id
  key string
  key = localhost + UTC()
 
  newGroupMember GroupMember
  newGroupMember = NewGroupMember(key,localhost,0)
  
  /* 
    use this to test (from command line for now):
      echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon, err := udp.NewDaemon(udpHost)
  if err != nil {
    log.Panic("Daemon creation", err)
  }
   //Add current machine to the list - Index by 'Address + Current Time' - will be unique even aftr machine leaves - Set heartbeat to 0
  if groupMember != "" {
    //Send the groupMember this machines list
    daemon.Gossip(newGroupMember, groupMember)
    logger.Log("GOSSIP","Gossiping new member to the group")
  }
  daemon.ReceiveDatagrams()
  log.Println("Done. (shouldn't be here)") // Cant make sense of logline comment either

  // Generate an ID to join the group with (may not be possible)
  // Join the group, receive current group list (with self included) - - Not sure where you receive this ? Receive Datagrams - Where is that stored

    ///@khshah3 - You seem to not have committed the udp or a lot of the code you had already done so I am a bit confused as to what is going so I am just trying to fill in where I can atleast with the logic.


    //Assuming this is the recieved grouplist
    receivedGroupList Map[String]GroupMember
    
    // LOOP for every heartbeat:
    //Should probably do this asynchronously - Duration set ticker
    ticker := time.Tick(duration * time.Second)
    for now := range ticker {
    


    
    // golang language spec explicitly defines maps as having undefined ordering of keys. Furthermore, since Go 1, key order is intentionally randomized between runs to prevent dependency          
    // on any perceived order. i.e Take the first one O(1)
    
    gossipMemberKey string
    index int
    index = 0
    //Store the first groupMemberKey which will be random
    for groupMemberKey, value := range groupList {
        if (index == 0 )
            gossipMemberKey = groupMemberKey
        value.IncrementHeartBeat
        index++
    }
    
    //Set current heartbeat to 0 - Do it after iteration to ensure self goes to 0
    groupList[localhost].SetHeartBeat(0)
    
    //Gossip the current list to random member
    daemon.Gossip(groupList,groupList[groupMemberKey].address)
        
            
    }


    
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




