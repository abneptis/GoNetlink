package netlink

import "strconv"

// Implements net.Addr for netlink addresses.
type Address struct {
  pid int
}

// Returns "netlink"
func (self Address)Network()(string){ return "netlink" }
// Ewruena the "pid" of the request.
func (self Address)Address()(string){ return strconv.Itoa(self.pid) }
