package data

import (
)

/*
  Represents the status of a single machine in the group
*/

type GroupMember struct {
  Id, Address string
  Heartbeat int
}


