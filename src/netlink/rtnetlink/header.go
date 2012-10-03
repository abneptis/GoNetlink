package rtnetlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

// An RTNetlink Header describes the structures
// between the netlink header and the nlattributes.
type Header interface {
	Len() int // The (unpadded) length of the Header.
	UnmarshalNetlink([]byte, int) error
	MarshalNetlink(int) ([]byte, error)
}
