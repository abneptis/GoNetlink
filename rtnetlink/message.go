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


func (self Message)MarshalNetlink(pad int)(out []byte, err os.Error){
  buff := bytes.NewBuffer(nil)
  bb, err := self.Header.MarshalNetlink(pad)
  if err == nil {
    buff.Write(bb)
    bb, err = netlink.MarshalAttributes(self.Attributes, pad)
    if err == nil {
      buff.Write(bb)
    }
  }
  out = netlink.PadBytes(buff.Bytes(), pad)
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
