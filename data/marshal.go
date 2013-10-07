package data

import (
)

/*
  This will be responsible for the data conversion. It allows us to take an object like 
  GroupMember and change it to an array of bytes to send over UDP. Then, on the other end,
  it should convert these bytes back into the original object. 
*/


// The following are just examples to help think about the problem
// The final implementation may look nothing like this

type MyMarshal struct {
}

// Serialize a GroupMember for transmission over UDP
func Marshal(member *GroupMember) (serialized string) {
  serialized = "TODO, this should be a groupMember's heartbeats"
  return
}

// Deserialize a transmitted GroupMember
func Unmarshal(serialized string) (member *GroupMember) {
  return
}
