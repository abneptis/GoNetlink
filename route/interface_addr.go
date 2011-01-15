package route

import "netlink"
import "os"
import "bytes"
//import "log"

type InterfaceAddressMessage struct {
  Family byte
  Prefix byte
  Flags byte
  Scope LinkScope
  Index int32
  Attributes []netlink.Attribute
}

func (self InterfaceAddressMessage)GetAttribute(t netlink.AttributeType)(attr netlink.Attribute, err os.Error){
  for i := range(self.Attributes){
    if self.Attributes[i].AttributeType() == t {
      attr = self.Attributes[i]
      break
    }
  }
  if attr == nil {
    err = os.NewError("Attribute not found")
  }
  return
}

func NewInterfaceAddrMesage(f, p, fl byte, s LinkScope, idx int32,  attrs []netlink.Attribute)(*InterfaceAddressMessage){
  return &InterfaceAddressMessage{
    Family: f,
    Prefix: p,
    Flags: fl,
    Scope: s,
    Index: idx,
    Attributes: attrs,
  }
}

func (self InterfaceAddressMessage)Marshal()(out []byte, err os.Error){
  buf := bytes.NewBuffer(nil)
  enc := netlink.NewEncoder(buf)
  err = enc.Encode([]byte{self.Family, self.Prefix, self.Flags, byte(self.Scope)})
  if err != nil { return }
  err = enc.Encode(self.Index)
  if err != nil { return }
  for ai := range(self.Attributes) {
    var bb []byte
    bb, err = netlink.Marshal(self.Attributes[ai])
    if err != nil { break }
    err = enc.Encode(uint16(4 + len(bb)))
    if err != nil { break }
    err = enc.Encode(self.Attributes[ai].AttributeType())
    if err != nil { break }
    err = enc.Encode(bb)
    if err != nil { break }
    //log.Printf("Padding attribute %d (%d)[%d: %X]", ai, self.Attributes[ai].AttributeType(), len(bb), bb)
    switch len(bb) % 4 {
        case 0:
        case 1: err = enc.Encode([]byte{0,0,0})
        case 2: err = enc.Encode([]byte{0,0})
        case 3: err = enc.Encode([]byte{0})
    }
    if err != nil { break }
  }
  if err == nil {
    out = buf.Bytes()
  }
  return
}

func (self *InterfaceAddressMessage)Unmarshal(in []byte)(err os.Error){
  if len(in) < 8 {
    return os.NewError("Message too short to be a valid RTNETLINK/IAM")
  }
  err = netlink.Unmarshal(in[0:1], &self.Family)
  if err != nil { return }
  // TODO: Check padding?
  err = netlink.Unmarshal(in[1:2], &self.Prefix)
  if err != nil { return }
  err = netlink.Unmarshal(in[2:3], &self.Flags)
  if err != nil { return }
  err = netlink.Unmarshal(in[3:4], &self.Scope)
  if err != nil { return }
  err = netlink.Unmarshal(in[4:8], &self.Index)
  if err != nil { return }
  pos := 8
  for pos < len(in) {
    var attrlen uint16
    err = netlink.Unmarshal(in[pos:pos+2], &attrlen)
    if err != nil { break }
    var attr netlink.Attribute
    attr, err = netlink.UnmarshalAttribute(in[pos+2:pos+int(attrlen)])
    if err != nil { break }
    //log.Printf("\tUnmarshalled Attribute: ATTR: 0x%X [%X]", attr.AttributeType(), attr.Body())
    pos += int(attrlen)
    self.Attributes = append(self.Attributes, attr)
    switch pos % 4 {
      case 0:
      default:
        pos += 4 - (pos % 4)
    }
  }
  return
}
