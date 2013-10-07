package data

import (
  "fmt"
  //"log"
  "strings"
  "strconv"
)
/*
  This will be responsible for the data conversion. It allows us to take an object like 
  GroupMember and change it to an array of bytes to send over UDP. Then, on the other end,
  it should convert these bytes back into the original object. 
*/

const (
  delim = "$$$"
)

// Serialize a GroupMember for transmission over UDP
func Marshal(member *GroupMember) (serialized string) {
  if member == nil {
    return "NIL"
  }

  serialized = fmt.Sprintf("%s%s%d", member.Id, delim, member.Heartbeat)
  //log.Printf("<%s, %d> ---> %s", member.Id, member.Heartbeat, serialized)
  return
}

// Deserialize a transmitted GroupMember
func Unmarshal(serialized string) (member *GroupMember) {
  if serialized == "NIL" {
    return nil
  }

  fields := strings.SplitN(serialized, delim, 2)
  id, hbs := fields[0], fields[1]
  address := strings.SplitN(id, "###", 2)[0]
  hb, _ := strconv.Atoi(hbs)
  return NewGroupMember(id, address, hb)
}
