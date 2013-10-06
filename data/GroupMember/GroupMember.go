package GroupMember

// TODO : There is a different GroupMember committed, we need to decide which one to tuse
import (
    "log"
    "fmt"
)

type GroupMember struct {
     machineId string
     ipAddress string
     heartBeat int
}

//Creates a new group member
func (member *GroupMember) NewGroupMember(machineId string , ipAddress string, heartBeat int){
    member.machineId = machineId
    member.ipAddress = ipAddress
    member.heartBeat = heartBeat
    log.Println("INFO","Creating a new group member of " + machineId + ipAddress)
}

//Increments the heartbeat of the current group member
func (member *GroupMember) IncrementHeartBeat() {
  log.Println("INFO", fmt.Sprintf("Incrementing Heart Beat of Machine %s current heartbeat: %i", member.machineId, member.heartBeat))
    member.heartBeat++
}

//Sets the heartbeat of the current group member
func (member *GroupMember) SetHeartBeat(heartbeat int)(){
    log.Println("INFO",fmt.Sprintf("Setting heartbeat of machine %s current hearbeat: %i new heartbeat: %i", member.machineId, member.heartBeat, heartbeat))
    member.heartBeat = heartbeat
}


