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

func _pad(in []byte, pad int)(out []byte){
  if pad > 0 {
    pblk := (len(in) + 1) / pad
    fsize := pblk * pad
    if fsize != len(in) {
      out = make([]byte, fsize)
      copy(out, in)
    } else {
      out = in
    }
  }
  return
}

func _repos(pos int, pad int)(out int){
  if pad > 0 {
    out = pad * ((pos + (pad - 1) ) / pad)
  } else {
    out = pos
  }
  return
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
  out = _pad(buff.Bytes(), pad)
  return
}

func (self *Message)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) < self.Header.Len() {
    return os.NewError("Insufficient data for unmarshal of Header")
  }
  err = self.Header.UnmarshalNetlink(in[0:self.Header.Len()], pad)
  if err == nil {
    pos := _repos(self.Header.Len(), pad)
    if len(in) > pos {
      self.Attributes, err = netlink.UnmarshalAttributes(in[pos:], pad)
    }
  }
  return
}
