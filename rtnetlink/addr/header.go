package addr

import "encoding/binary"
import "netlink/rtnetlink"
import "os"


const HEADER_LENGTH = 8
type Header [HEADER_LENGTH]byte

func NewHeader(afam byte, pl uint8, fl byte, scope rtnetlink.Scope, ifindex uint32)(*Header){
  hdr := Header{afam, pl, fl, byte(scope)}
  binary.LittleEndian.PutUint32(hdr[4:8], ifindex)
  return &hdr
}


func (self Header)Len()(int){ return HEADER_LENGTH }
func (self Header)AddressFamily()(byte){ return self[0] }
func (self Header)PrefixLength()(uint8){ return self[1] }
func (self Header)Flags()(uint8){ return self[2] }
func (self Header)Scope()(rtnetlink.Scope){ return rtnetlink.Scope(self[3]) }
func (self Header)InterfaceIndex()(uint32){ return binary.LittleEndian.Uint32(self[4:8]) }

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
