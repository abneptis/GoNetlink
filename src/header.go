package netlink

import "os"
import "encoding/binary"

const HEADER_LENGTH = 16

// Represents the header of a netlink.Message
type Header [HEADER_LENGTH]byte

func (self Header)Len()(int){ return HEADER_LENGTH}

func (self Header)MessageLength()(uint32) { return binary.LittleEndian.Uint32(self[0:4]) }
func (self *Header)SetMessageLength(in uint32) {  binary.LittleEndian.PutUint32(self[0:4], in)}
func (self Header)MessageType()(MessageType) { return MessageType(binary.LittleEndian.Uint16(self[4:6]))}
func (self Header)MessageFlags()(MessageFlags) { return MessageFlags(binary.LittleEndian.Uint16(self[6:8]))}
func (self Header)MessageSequence()(uint32) { return binary.LittleEndian.Uint32(self[8:12])}
func (self *Header)SetMessageSequence(in uint32) {  binary.LittleEndian.PutUint32(self[8:12], in)}
func (self Header)MessagePid()(uint32) { return binary.LittleEndian.Uint32(self[12:16])}

func NewHeader(t MessageType, f MessageFlags, seq uint32)(h *Header){
  h = &Header{}
  binary.LittleEndian.PutUint32(h[0:4], HEADER_LENGTH)
  binary.LittleEndian.PutUint16(h[4:6], uint16(t))
  binary.LittleEndian.PutUint16(h[6:8], uint16(f))
  binary.LittleEndian.PutUint32(h[8:12], seq)
  binary.LittleEndian.PutUint32(h[12:16], 0)

  return
}

func (self Header)MarshalNetlink(pad int)(out []byte, err os.Error){
  out = make([]byte, HEADER_LENGTH)
  copy(out, self[0:HEADER_LENGTH])
  out = PadBytes(out, pad)
  return
}

func (self *Header)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) !=  HEADER_LENGTH {
    err = os.NewError("Incorrect NetlinkHeader length")
  } else {
    copy(self[0:HEADER_LENGTH], in[0:HEADER_LENGTH])
  }
  return
}
