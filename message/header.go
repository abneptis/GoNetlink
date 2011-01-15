package netlink

type header struct {
  Length uint32
  Type  MessageType
  Flags MessageFlags
  Sequence uint32
  Pid      uint32
}

func NewHeader(t MessageType, f MessageFlags, seq uint32)(*header){
  return &header{Type:t, Flags: f, Sequence: seq}
}
