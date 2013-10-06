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
func (self *MyMarshal) Marshal(member *GroupMember) (serialized []byte) {
  return
}

// Deserialize a transmitted GroupMember
func (self *MyMarshal) Unmarshal(serialized []byte) (member *GroupMember) {
  return
}
