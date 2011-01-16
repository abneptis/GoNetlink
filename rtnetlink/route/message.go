package route

import "encoding/binary"
import "netlink"
import "netlink/rtnetlink"
import "os"


type Message struct {
  header [12]byte
  attributes []netlink.Attribute
}

func NewMessage(afam byte, dl uint8, sl uint8, tos uint8, t Table, o Origin, s rtnetlink.Scope, T Type, f Flags)(*Message){
  hdr := [12]byte{afam, dl, sl, tos, byte(t), byte(o), byte(s), byte(T)}
  binary.LittleEndian.PutUint32(hdr[8:12], uint32(f))
  return &Message{
    header: hdr,
  }
}

func (self Message)AddressFamily()(byte){ return self.header[0] }
func (self Message)AddressDestLength()(uint8){ return self.header[1] }
func (self Message)AddressSourceLength()(uint8){ return self.header[2] }
func (self Message)TOS()(uint8){ return self.header[3] }
func (self Message)RoutingTable()(Table){ return Table(self.header[4]) }
func (self Message)RouteOrigin()(Origin){ return Origin(self.header[5]) }
func (self Message)AddressScope()(rtnetlink.Scope){ return rtnetlink.Scope(self.header[6]) }
func (self Message)RouteType()(Type){ return Type(self.header[7]) }
func (self Message)Flags()(Flags) { return Flags(binary.LittleEndian.Uint32(self.header[8:12])) }

func (self Message)Attributes()([]netlink.Attribute) { return self.attributes }
func (self *Message)AddAttribute(a netlink.Attribute)(){ self.attributes = append(self.attributes, a) }

func (self *Message)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) < 12 {
    err = os.NewError("Too short to be a valid Routing Message")
  }
  if err == nil {
    copy(self.header[0:12], in[0:12])
    self.attributes, err = netlink.UnmarshalAttributes(in[12:], pad)
  }
  return
}

func (self Message)MarshalNetlink(pad int)(out []byte, err os.Error){
  attrbytes, err := netlink.MarshalAttributes(self.attributes, pad)
  if err == nil {
    out = make([]byte, len(attrbytes) + 12)
    copy(out, self.header[0:])
    copy(out[12:], attrbytes)
  }
  return
}
