package netlink

import "os"

type MessageFlags uint16

// include/linux/netlink.h
// ALL
const (
  NLM_F_REQUEST MessageFlags = 1 << iota
  NLM_F_MULTI
  NLM_F_ACK
  NLM_F_ECHO
)

// GET
const (
  NLM_F_ROOT    MessageFlags = 0x100 << iota
  NLM_F_MATCH
  NLM_F_ATOMIC
  NLM_F_DUMP MessageFlags = (NLM_F_ROOT|NLM_F_MATCH)
)

// NEW
const (
  NLM_F_REPLACE  MessageFlags =  0x100 << iota
  NLM_F_EXCL
  NLM_F_CREATE
  NLM_F_APPEND
)

func (self MessageFlags)Marshal()([]byte, os.Error){
  return Marshal(uint16(self))
}

func (self *MessageFlags)Unmarshal(in []byte)( err os.Error){
  var mt uint16
  err = Unmarshal(in, &mt)
  if err == nil { *self = MessageFlags(mt) }
  return
}

