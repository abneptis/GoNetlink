package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "fmt"
import "bytes"
import "os"
import "io"

// A netlink message contains a Netlink header,
// and a body of bytes.
type Message struct {
  Header *Header
  Body []byte
}

// NetlinkMarshaler's are used to pad and format netlink data.
type NetlinkMarshaler interface {
  MarshalNetlink(int)([]byte, os.Error)
}

// Creates a new message from a marshalable object
func NewMessage(t MessageType, f MessageFlags, u NetlinkMarshaler, pad int)(msg *Message, err os.Error){
  msg = &Message{Header: NewHeader(t,f,0) }
  msg.Body, err = u.MarshalNetlink(pad)
  if err == nil {
    msg.Header.SetMessageLength(uint32(msg.Header.Len()) + uint32(len (msg.Body)))
  }
  return
}

// Reads a message from an io.Reader, with a specified padding.
// NB: Netlink uses a very strict protocol, and it is encouraged
// that r be a bufio.Reader
func ReadMessage(r io.Reader, pad int)(msg *Message, err os.Error){
  var n int
  msg = &Message{Header: &Header{}}
  ib := make([]byte, msg.Header.Len())
  n, err = r.Read(ib)
  if err == nil && n != msg.Header.Len() {
    err = os.NewError("Incomplete netlink header")
  }
  if err == nil {
    err = msg.Header.UnmarshalNetlink(ib, pad)
    if err == nil {
      msg.Body = make([]byte, msg.Header.MessageLength() - uint32(msg.Header.Len()))
      n, err = r.Read(msg.Body)
      if err == nil && n != int(msg.Header.MessageLength()) - int(msg.Header.Len()) {
        err = os.NewError(fmt.Sprintf("Incomplete netlink message body (Got: %d; Expected: %d)", n, int(msg.Header.MessageLength()) - int(msg.Header.Len())))
      } else {
      }
    }
  }
  return
}

// Marshals a message with appropriate padding (generally none).
func (self Message)MarshalNetlink(pad int)(out []byte, err os.Error){
  hdrout, err := self.Header.MarshalNetlink(pad)
  if err == nil {
    out = bytes.Join([][]byte{hdrout, self.Body}, []byte{})
  }
  return
}
