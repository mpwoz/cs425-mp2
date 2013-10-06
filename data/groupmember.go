package data

import (
  "log"
  "fmt"
)

/*
  Represents the status of a single machine in the group
*/

type GroupMember struct {
  Id, Address string
  Heartbeat int
}

// Initialize a new group member
func NewGroupMember(machineId, address string, heartBeat int) (member *GroupMember) {
  log.Println("INFO", fmt.Sprintf("Creating a new group member of %s, %s", machineId, address))
  member = new(GroupMember)
  member.Id = machineId
  member.Address = address
  member.Heartbeat = heartBeat
  return
}

func (self *GroupMember) IncrementHeartBeat() {
  log.Println("INFO", fmt.Sprintf("Incrementing Heart Beat of Machine %s current heartbeat: %i", self.Id, self.Heartbeat))
  self.Heartbeat++
}

func (self *GroupMember) SetHeartBeat(heartbeat int) {
    log.Println("INFO",fmt.Sprintf("Setting heartbeat of machine %s current hearbeat: %i new heartbeat: %i", self.Id, self.Heartbeat, heartbeat))
    self.Heartbeat = heartbeat
}


