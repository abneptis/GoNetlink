package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "strconv"

// Implements net.Addr for netlink addresses.
type Address struct {
  pid int
}

// Returns "netlink"
func (self Address)Network()(string){ return "netlink" }
// Ewruena the "pid" of the request.
func (self Address)Address()(string){ return strconv.Itoa(self.pid) }
