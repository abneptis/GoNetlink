package route

import "encoding/binary"
import "netlink"
import "os"


type RoutingMessage struct {
  header [12]byte
  attributes []netlink.Attribute
}

func (self RoutingMessage)AddressFamily()(byte){ return self.header[0] }
func (self RoutingMessage)AddressDestLength()(uint8){ return self.header[1] }
func (self RoutingMessage)AddressSourceLength()(uint8){ return self.header[2] }
func (self RoutingMessage)TOS()(uint8){ return self.header[3] }
func (self RoutingMessage)RoutingTable()(Table){ return Table(self.header[4]) }
func (self RoutingMessage)RouteOrigin()(Origin){ return Origin(self.header[5]) }
func (self RoutingMessage)AddressScope()(Scope){ return Scope(self.header[6]) }
func (self RoutingMessage)RouteType()(Type){ return Type(self.header[7]) }
func (self RoutingMessage)Flags()(Flags) { return Flags(binary.LittleEndian.Uint32(self.header[8:12])) }

func (self *RoutingMessage)UnmarshalNetlink(in []byte)(err os.Error){
  if len(in) < 12 {
    err = os.NewError("Too short to be a valid Routing Message")
  }
  if err == nil {
    copy(self.header[0:12], in[0:12])
    self.attributes, err = netlink.UnmarshalAttributes(in[12:], 4)
  }
  return
}

func (self *RoutingMessage)MarshalNetlink()(out []byte, err os.Error){
  attrbytes, err := netlink.MarshalAttributes(self.attributes, 4)
  if err == nil {
    out = make([]byte, len(attrbytes) + 12)
    copy(out, self.header[0:])
    copy(out[12:], attrbytes)
  }
  return
}
