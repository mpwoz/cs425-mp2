package GroupMember

import (
    "log"
    "os"
    "bufio"
  "../logger"
)

type GroupMember struct {
     machineId string
     ipAddress string
     heartBeat int
}

//Creates a new group member
func (member *GroupMember) NewGroupMember(machineId string , ipAddress string, heartBeat ,){
    member.machineId = machineId
    member.ipAddress = ipAddress
    member.heartBeat = heartbeat
    logger.Log ("INFO","Creating a new group member of " + machineId + ipAddress)
    
}

//Increments the heartbeat of the current group member
func (member *GroupMember)IncrementHeartBeat()(int){
    logger.Log("INFO","Incrementing Heart Beat of Machine" + member.machineId + "current hearbeat:" + member.heartBeat)
    ++member.heartbeat
}

//Sets the heartbeat of the current group member
func (member *GroupMember) SetHeartBeat(heartbeat int)(){
    logger.Log("INFO","Setting heartbeat of machine" + member.machineId + "current hearbeat:" + member.heartBeat + "new heartbeat:" + heartbeat)
    member.heartbeat = heartbeat
}


