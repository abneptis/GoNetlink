package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "fmt"

import "bytes"
import (
	"encoding/binary"
	"errors"
)

// A basic netlink type used for identifying
// nlattrs in a message.
type AttributeType uint16

// An attribute is used to hold an Netlink Attribute.
// An attribute is stored as a Length-Type-Value tuple.
// Length and Type are 16 bit integers, so values may not
// exceed 2^16.
type Attribute struct {
	Type AttributeType
	Body []byte
}

// Marshals a netlink attribute as a full LTV tuple.
func (self Attribute) MarshalNetlink(pad int) (out []byte, err error) {
	l := len(self.Body)
	out = make([]byte, l+4)
	binary.LittleEndian.PutUint16(out[0:2], uint16(len(self.Body)+4))
	binary.LittleEndian.PutUint16(out[2:4], uint16(self.Type))
	copy(out[4:], self.Body[0:])
	out = PadBytes(out, pad)
	return
}

// Unmarshals a netlink attribute.
func UnmarshalAttributes(in []byte, padding int) (out []Attribute, err error) {
	pos := 0
	for pos < len(in) {
		l := binary.LittleEndian.Uint16(in[pos : pos+2])
		if int(l) > len(in)-pos {
			err = errors.New("Can't parse attribute (too long)")
			break
		}
		if l > 4 {
			t := binary.LittleEndian.Uint16(in[pos+2 : pos+4])
			out = append(out, Attribute{Type: AttributeType(t), Body: in[pos+4 : pos+int(l)]})
			pos = Reposition(pos+int(l), padding)
		} else {
			err = errors.New(fmt.Sprintf("Invalid Attributeibute (Len: %d):", l))
			break
		}
	}
	return
}

// Returns the padded bytes of a marshalled list of attributes.
// Any marshalling error will cause the sequence to abort.
func MarshalAttributes(in []Attribute, padding int) (out []byte, err error) {
	for i := range in {
		var b []byte
		b, err = in[i].MarshalNetlink(padding)
		if err == nil {
			out = bytes.Join([][]byte{out, b}, []byte{})
		} else {
			break
		}
	}
	return
}
