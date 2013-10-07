package data

import (
/*
    //"bytes"
    //"encoding/binary"
    //"fmt"
=======*/
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
/*<<<<<<< HEAD
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
=======*/
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
