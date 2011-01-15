package netlink

//import "log"
import "os"
import "bytes"

type NetlinkMessage interface {
  //NetlinkLen()(uint32) // May NOT match the actual size of the data due to padding!
  MessageType()(MessageType)
  MessageFlags()(MessageFlags)
  Sequence()(uint32)
  SetSequence(uint32)
  Pid()(uint32)
  Body()([]byte)
}

type nlmsg struct {
  _Type MessageType
  _Flags MessageFlags
  _Sequence uint32
  _Pid uint32
  _Body []byte
}

func NewMessage(t MessageType, f MessageFlags, umsg Marshaler)(msg *nlmsg, err os.Error){
  msg = &nlmsg{_Type:t, _Flags: f, _Sequence: 0, _Pid: 0}
  msg._Body, err = umsg.Marshal()
  return
}

func(self *nlmsg)MessageType()(MessageType){ return self._Type }
func(self *nlmsg)MessageFlags()(MessageFlags){ return self._Flags }
func(self *nlmsg)Sequence()(uint32){ return self._Sequence }
func(self *nlmsg)SetSequence(seq uint32){ self._Sequence = seq }
func(self *nlmsg)Pid()(uint32){ return self._Pid }
func(self *nlmsg)Body()([]byte){ return self._Body}

func (self *nlmsg)Marshal()(out []byte, err os.Error){
  buff := bytes.NewBuffer(nil)
  enc := NewEncoder(buff)
  err = enc.Encode(uint32(len(self._Body) + 16)) // Body + NL Header (Plus sizeof size).
  if err != nil { return }
  err = enc.Encode(self._Type)
  if err != nil { return }
  err = enc.Encode(self._Flags)
  if err != nil { return }
  err = enc.Encode(self._Sequence)
  if err != nil { return }
  err = enc.Encode(self._Pid)
  if err != nil { return }
  err = enc.Encode(self._Body)
  if err != nil { return }
  out = buff.Bytes()
  //log.Printf("Marshal:: Body [%X]", self._Body)
  //log.Printf("Marshal:: Out [%X]", out)
  return
}


// Note, length has already been stripped off of a message at this point!
func (self *nlmsg)Unmarshal(in []byte)(err os.Error){
  err = Unmarshal(in[0:2], &self._Type)
  if err != nil { return }
  err = Unmarshal(in[2:4], &self._Flags)
  if err != nil { return }
  err = Unmarshal(in[4:8], &self._Sequence)
  if err != nil { return }
  err = Unmarshal(in[8:12], &self._Pid)
  if err != nil { return }
  err = Unmarshal(in[12:], &self._Body)
  //log.Printf("Body: %X", self._Body)
  return
}
