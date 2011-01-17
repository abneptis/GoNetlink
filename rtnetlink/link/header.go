package link

import "os"
import "encoding/binary"


const HEADER_LENGTH = 16
type Header [16]byte

func (self Header)Len()(int) { return HEADER_LENGTH }
func (self *Header)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) != HEADER_LENGTH {
    err = os.NewError("Wrong length for Header")
  } else {
    copy(self[0:HEADER_LENGTH], in[0:HEADER_LENGTH])
  }
  return
}

func (self Header)MarshalNetlink(pad int)(out []byte, err os.Error){
  out = self[0:HEADER_LENGTH]
  return
}


func (self Header)InterfaceFamily()(byte){ return self[0]}
func (self Header)InterfaceType()(uint16){ return binary.LittleEndian.Uint16(self[2:4]) }
func (self Header)InterfaceIndex()(uint32){ return binary.LittleEndian.Uint32(self[4:8]) }
func (self Header)InterfaceFlags()(uint32){ return binary.LittleEndian.Uint32(self[8:12]) }
func (self Header)InterfaceChanges()(uint32){ return binary.LittleEndian.Uint32(self[12:16]) }

