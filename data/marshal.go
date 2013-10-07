package data

import (
    "bytes"
    "encoding/binary"
    "fmt"
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
func Marshal(member *GroupMember) (serialized []byte) {
  
  var Id [40]byte
  var Address [120]byte
  
  
  copy(Id[:], member.Id)
  copy(Address[:], member.Address)
  Heartbeat := int8(member.Heartbeat)
  
  IdLength := byte(len(member.Id))
  AddressLength := byte(len(member.Address))
  fmt.Println(IdLength)
  fmt.Println(AddressLength)
  fmt.Println(len(member.Address))
  fmt.Println(Heartbeat)
  
  buf := new(bytes.Buffer)
  err := binary.Write(buf, binary.LittleEndian, IdLength)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
  err = binary.Write(buf, binary.LittleEndian, Id)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
  err = binary.Write(buf, binary.LittleEndian, AddressLength)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
       err = binary.Write(buf, binary.LittleEndian, Address)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
      err = binary.Write(buf, binary.LittleEndian, Heartbeat)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
  return buf.Bytes()
}

// Deserialize a transmitted GroupMember
func UnMarshal(serialized []byte) (member *GroupMember) {
  
  var machineId string
  var address string
  var Heartbeat int8
  
  var Id [40]byte
  var Address [120]byte
  
  // 
  buf := bytes.NewBuffer(serialized)
  var IdLength int8
  var AddressLength int8
  err := binary.Read(buf, binary.LittleEndian, &IdLength)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
     err = binary.Read(buf, binary.LittleEndian, &Id)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
    fmt.Println(IdLength)
    machineId = string(Id[:IdLength])
    
     err = binary.Read(buf, binary.LittleEndian, &AddressLength)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
     err = binary.Read(buf, binary.LittleEndian, &Address)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
    fmt.Println(AddressLength)
    
    address = string(Address[:AddressLength])
    
     err = binary.Read(buf, binary.LittleEndian, &Heartbeat)
	if err != nil {
		fmt.Println("binary.Read failed:", err)
	}
    
    fmt.Println(machineId)
    fmt.Println(address)
    fmt.Println(Heartbeat)
    member = NewGroupMember(machineId, address, int(Heartbeat))
    return
}
