package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "log"

import (
	"encoding/binary"
	"errors"
)

import "bytes"

// The attributeFinder interface  should be used by READERS,
// Writers will need to use an AttributeWriter
type AttributeFinder interface {
	GetAttribute(AttributeType) (Attribute, error)
}

// AttributeWriters allow attributes to be added/updated.
type AttributeWriter interface {
	SetAttribute(Attribute)
}

//  Sets the attribute value to the string specified by 's'.
// s will be NULL terminated.
func SetAttributeCString(aw AttributeWriter, at AttributeType, s string) {
	buff := bytes.NewBufferString(s + "\x00")
	log.Printf("Buff: %X", buff.Bytes())
	aw.SetAttribute(Attribute{Type: at, Body: buff.Bytes()})
	return
}

// Returns the attribute as an uint32 (or an error if the attr size is not
// 32 bits.
func GetAttributeUint32(af AttributeFinder, at AttributeType) (out uint32, err error) {
	attr, err := af.GetAttribute(at)
	if err == nil {
		body := attr.Body
		if len(body) != 4 {
			err = errors.New("Attribute wrong size for Uint32")
		} else {
			out = binary.LittleEndian.Uint32(body)
		}
	}
	return
}

// Gets an attribute value as a string.
// Note, for much of RTNetlink, you want GetAttributeCString,
// which will verify and chop the  tailing NULL.
func GetAttributeString(af AttributeFinder, at AttributeType) (out string, err error) {
	attr, err := af.GetAttribute(at)
	if err == nil {
		out = string(attr.Body)
	}
	return
}

// Same as GetAttributeString, but expects the string to be NULL terminated,
// the null terminator will be stripped..
func GetAttributeCString(af AttributeFinder, at AttributeType) (out string, err error) {
	attr, err := af.GetAttribute(at)
	if err == nil {
		outbody := attr.Body
		if len(outbody) == 0 {
			err = errors.New("Invalid body")
		} else {
			if outbody[len(outbody)-1] != 0 {
				err = errors.New("Expected NULL-terminated string")
			} else {
				out = string(outbody[0 : len(outbody)-1])
			}
		}
	}
	return
}
