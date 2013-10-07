package main

import (
  "./data"
  "./udp"
  "flag"
  "log"
  "time"
  "math/rand"
  "fmt"
)

//Get random member & Increment HeartBeat
func getRandomMemberAndIncrementAll(groupList map[string]data.GroupMember)(data.GroupMember){
    randomMemberIndex := rand.Int() % len(groupList)
    fmt.Println(randomMemberIndex)
    index := 0
    var randomMember data.GroupMember
    for _, currMember := range groupList {
    
      if(randomMemberIndex == index){
            randomMember = currMember
            fmt.Println("Random Member Found")
        }
      currMember.IncrementHeartBeat()
      index++
      
    }
    return randomMember

}
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

  groupList := make(map[string]data.GroupMember)
  
  newMember := data.NewGroupMember("KEY", "VALUE", 3);
  groupList["KEY"] = *newMember
  
  // Determine the heartbeat duration time
  duration := 5

  /* 
    use this to test (from command line for now):
      echo -n "hello" >/dev/udp/localhost/4567
  */
  daemon, err := udp.NewDaemon(udpHost)
  if err != nil {
    log.Panic("Daemon creation", err)
  }

  // Join the group by broadcasting a dummy message
  if groupMember != "" {
    udp.SendMessage("JOIN", groupMember)
    log.Println("GOSSIP","Gossiping new member to the group") // TODO use mp1 logger
  }
    

  // LOOP for every heartbeat:
  //Should probably do this asynchronously - Duration set ticker
  
  ticker := time.Tick(time.Duration(duration) * time.Second)
  
//TODO: Figure out how to move list to Daemon and use the ticker and listen to UDP in parallel
  for now := range ticker {
  
    //Get random member , increment current members
    randomMember := getRandomMemberAndIncrementAll(groupList)
    log.Println("TICK",now)
    
    sendingData := data.Marshal(&randomMember)
    fmt.Println(sendingData)
    checkMember := data.UnMarshal(sendingData)
    
    fmt.Println(checkMember)
    //TODO Send Marshalled data
    
    udp.SendMessage("sendingData",randomMember.Address)
    
    
daemon.ReceiveDatagrams()
  }
 
  // Broadcast current group state to chosen machine
  // Increment heartbeat counters for all machines (zero own counter)
}



