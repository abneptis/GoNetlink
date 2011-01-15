package rtnetlink

import "netlink"
import "os"
import "bytes"

type InterfaceInformationMessage struct {
  Family byte
  // 1 byte padding
  Type uint16
  Index int32
  Flags InterfaceFlags
  Changes InterfaceFlags
  Attributes []netlink.Attribute
}

func (self InterfaceInformationMessage)GetAttribute(t netlink.AttributeType)(attr netlink.Attribute, err os.Error){
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

func NewInterfaceInfoMesage(f byte, dtype uint16, ifindex int32, flags InterfaceFlags, attrs []netlink.Attribute)(*InterfaceInformationMessage){
  return &InterfaceInformationMessage{
    Family: f,
    Type: dtype,
    Index: ifindex,
    Flags: flags,
    Changes: InterfaceFlags(0xffffffff), // currently netlink expects this to be set to '-1'.
    Attributes: attrs,
  }
}

func (self InterfaceInformationMessage)Marshal()(out []byte, err os.Error){
  buf := bytes.NewBuffer(nil)
  enc := netlink.NewEncoder(buf)
  // Family + 1 byte of padding
  err = enc.Encode([]byte{self.Family, 0})
  if err != nil { return }
  err = enc.Encode(self.Type)
  if err != nil { return }
  err = enc.Encode(self.Index)
  if err != nil { return }
  err = enc.Encode(self.Flags)
  if err != nil { return }
  err = enc.Encode(self.Changes)
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

func (self *InterfaceInformationMessage)Unmarshal(in []byte)(err os.Error){
  if len(in) < 16 {
    return os.NewError("Message too short to be a valid RTNETLINK/IIM")
  }
  err = netlink.Unmarshal(in[0:1], &self.Family)
  if err != nil { return }
  // TODO: Check padding?
  err = netlink.Unmarshal(in[2:4], &self.Type)
  if err != nil { return }
  err = netlink.Unmarshal(in[4:8], &self.Index)
  if err != nil { return }
  err = netlink.Unmarshal(in[8:12], &self.Flags)
  if err != nil { return }
  err = netlink.Unmarshal(in[12:16], &self.Changes)
  if err != nil { return }
  pos := 16
  //log.Printf("ATTRBytes: {%X}", in[16:])
  for pos < len(in) {
    //log.Printf("Position: %v", pos)
    var attrlen uint16
    //log.Printf("\tUnmarshalling attrlen")
    err = netlink.Unmarshal(in[pos:pos+2], &attrlen)
    if err != nil { break }
    //log.Printf("\t attrlen: %v", attrlen)
    var attr netlink.Attribute
    attr, err = netlink.UnmarshalAttribute(in[pos+2:pos+int(attrlen)])
    if err != nil { break }
    //log.Printf("\tAttribute: ATTR: 0x%X [%X]", attr.AttributeType(), attr.Body())
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
