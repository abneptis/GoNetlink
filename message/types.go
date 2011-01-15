package netlink

import "os"

type MessageType uint16

// include/linux/netlink.h
const (
  NLMSG_UNSPECIFIED MessageType = iota
  NLMSG_NOOP
  NLMSG_ERROR
  NLMSG_DONE
  NLMSG_OVERRUN
  MIN_TYPE = 0x10
)

func (self MessageType)Marshal()([]byte, os.Error){
  return Marshal(uint16(self))
}

func (self *MessageType)Unmarshal(in []byte)( err os.Error){
  var mt uint16
  err = Unmarshal(in, &mt)
  if err == nil { *self = MessageType(mt) }
  return
}
