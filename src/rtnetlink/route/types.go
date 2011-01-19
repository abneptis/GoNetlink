package route
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

type Type byte

const (
  RTN_UNSPEC Type = iota
  RTN_UNICAST
  RTN_LOCAL
  RTN_BROADCAST
  RTN_ANYCAST
  RTN_MULTICAST
  RTN_BLACKHOLE
  RTN_UNREACHABLE
  RTN_PROHIBIT
  RTN_THROW
  RTN_NAT
  RTN_XRESOLVE
  __RTN_MAX
)

var TypeStrings = map[Type]string {
  RTN_UNSPEC: "RTN_UNSPEC",
  RTN_UNICAST: "RTN_UNICAST",
  RTN_LOCAL: "RTN_LOCAL",
  RTN_BROADCAST: "RTN_BROADCAST",
  RTN_ANYCAST: "RTN_ANYCAST",
  RTN_MULTICAST: "RTN_MULTICAST",
  RTN_BLACKHOLE: "RTN_BLACKHOLE",
  RTN_UNREACHABLE: "RTN_UNREACHABLE",
  RTN_PROHIBIT: "RTN_PROHIBIT",
  RTN_THROW: "RTN_THROW",
  RTN_NAT: "RTN_NAT",
  RTN_XRESOLVE: "RTN_XRESOLVE",
  __RTN_MAX: "__RTN_MAX",
}

func (self Type)String()(string){
  return TypeStrings[self]
}

