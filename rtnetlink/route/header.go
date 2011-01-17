package route

import "encoding/binary"
import "netlink/rtnetlink"
import "os"


type Header [12]byte

func NewMessage(afam byte, dl uint8, sl uint8, tos uint8, t Table, o Origin, s rtnetlink.Scope, T Type, f Flags)(*Header){
  hdr := Header{afam, dl, sl, tos, byte(t), byte(o), byte(s), byte(T)}
  binary.LittleEndian.PutUint32(hdr[8:12], uint32(f))
  return &hdr
}

func (self Header)AddressFamily()(byte){ return self[0] }
func (self Header)AddressDestLength()(uint8){ return self[1] }
func (self Header)AddressSourceLength()(uint8){ return self[2] }
func (self Header)TOS()(uint8){ return self[3] }
func (self Header)RoutingTable()(Table){ return Table(self[4]) }
func (self Header)RouteOrigin()(Origin){ return Origin(self[5]) }
func (self Header)AddressScope()(rtnetlink.Scope){ return rtnetlink.Scope(self[6]) }
func (self Header)RouteType()(Type){ return Type(self[7]) }
func (self Header)Flags()(Flags) { return Flags(binary.LittleEndian.Uint32(self[8:12])) }


func (self *Header)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) < 12 {
    err = os.NewError("Too short to be a valid Routing Message")
  }
  if err == nil {
    copy(self[0:12], in[0:12])
  }
  return
}

func (self Header)MarshalNetlink(pad int)(out []byte, err os.Error){
  if err == nil {
    out = make([]byte, 12)
    copy(out, self[0:])
    out = rtnetlink.PadBytes(out, pad)
  }
  return
}
