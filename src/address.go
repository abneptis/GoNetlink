package netlink

import "strconv"

type Address struct {
  pid int
}

func (self Address)Network()(string){ return "netlink" }
func (self Address)Address()(string){ return strconv.Itoa(self.pid) }
