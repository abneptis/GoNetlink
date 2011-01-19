package rtnetlink

import "bytes"
import "netlink"
import "os"

type Message struct {
  Header Header
  Attributes []netlink.Attribute
}

func NewMessage(h Header, attrs []netlink.Attribute)(*Message){
  return &Message{Header:h, Attributes: attrs}
}

func (self *Message)SetAttribute(attr netlink.Attribute){
  t := attr.Type
  for i := range(self.Attributes){
    if t == self.Attributes[i].Type {
      self.Attributes[i] = attr
      return
    }
  }
  self.Attributes = append(self.Attributes, attr)
  return
}

func (self Message)GetAttribute(t netlink.AttributeType)(attr netlink.Attribute, err os.Error){
  for i := range(self.Attributes){
    if t == self.Attributes[i].Type {
      attr = self.Attributes[i]
      return
    }
  }
  err = os.NewError("Attribute not found")
  return
}

func (self Message)MarshalNetlink(pad int)(out []byte, err os.Error){
  hb, err := self.Header.MarshalNetlink(pad)
  if err == nil {
    var bb []byte
    bb, err = netlink.MarshalAttributes(self.Attributes, pad)
    if err == nil {
      out = bytes.Join([][]byte{ hb, bb }, []byte{} )
    }
  }
  return
}

func (self *Message)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) < self.Header.Len() {
    return os.NewError("Insufficient data for unmarshal of Header")
  }
  err = self.Header.UnmarshalNetlink(in[0:self.Header.Len()], pad)
  if err == nil {
    pos := netlink.Reposition(self.Header.Len(), pad)
    if len(in) > pos {
      self.Attributes, err = netlink.UnmarshalAttributes(in[pos:], pad)
    }
  }
  return
}
