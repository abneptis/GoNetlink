package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "fmt"
import (
	"encoding/binary"
	"errors"
)

//import "syscall"

// Unlike other headers, errors MAY be longer than the minimum length.
const ERROR_LENGTH = HEADER_LENGTH + 4

// Represents a netlink Error message.
type Error [ERROR_LENGTH]byte

// The error code (-errno) of the netlink message.
// 0 is used for netlink ACK's.
func (self Error) Code() int32 {
	return int32(binary.LittleEndian.Uint32(self[0:4]))
}

// Marshals an error to the wire.
func (self Error) MarshalNetlink(pad int) (out []byte, err error) {
	out = PadBytes(self[0:ERROR_LENGTH], pad)
	return
}

// Unmarshals an error from a netlink message.
func (self *Error) UnmarshalNetlink(in []byte, pad int) (err error) {
	if len(in) < ERROR_LENGTH {
		return errors.New(fmt.Sprintf("Invalid netlink error length: %d", len(in)))
	}
	copy(self[0:ERROR_LENGTH], in)
	return
}


// Implements os.Error by using err.Error(-Code()) 
func (self Error) Error() string {
//	return Errstr(int(-self.Code()))
	return "Error string unimplemented"	
}
