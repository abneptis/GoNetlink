package rtnetlink
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "bytes"
import "netlink"
import "os"

// A Message contains a Header object and a series of attributes.
// It is extracted from the Body of a netlink.Message
type Message struct {
  Header Header
  Attributes []netlink.Attribute
}

// Create a new rtnetlink.Message based off of a header an list of attributes
// (which may be nil or empty).
func NewMessage(h Header, attrs []netlink.Attribute)(*Message){
  return &Message{Header:h, Attributes: attrs}
}

// Replace or append the attribute with the AttributeType of
// attr
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

// Retrieve (the first) Attribute identified by Type,
// returning an error if not found.
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

// Handles the appropriate calls to marshal the Header and Attribute values,
// and return an appropriately padded result.
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

// Unmarshals a generic message using the header as a guide.
// An error will be returned if the header cannot unmarshal properly,
// or any attribute in the series failed.
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
