package rtnetlink

import "netlink"
import "os"

const (
  IFLA_UNSPEC netlink.AttributeType = iota
  IFLA_ADDRESS
  IFLA_BROADCAST
  IFLA_IFNAME
  IFLA_MTU
  IFLA_LINK
  IFLA_QDISC
  IFLA_STATS
  IFLA_COST
  IFLA_PRIORITY
  IFLA_MASTER
  IFLA_WIRELESS
  IFLA_PROTINFO
  IFLA_TXQLEN
  IFLA_MAP
  IFLA_WEIGHT
  IFLA_OPERSTATE
  IFLA_LINKMODE
  IFLA_LINKINFO
  IFLA_NET_NS_PID
  IFLA_IFALIAS
  IFLA_MAX
)

type LinkScope byte
func (self LinkScope)Marshal()([]byte, os.Error){
  return []byte{byte(self)}, nil
}
func (self *LinkScope)Unmarshal(in []byte)(err os.Error){
  if len(in) == 1 {
    *self = LinkScope(in[0])
  } else {
    err = os.NewError("Invalid unmarshal (too long)")
  }
  return
}


const (
  RT_SCOPE_UNIVERSE LinkScope = 0
  RT_SCOPE_SITE LinkScope = 200
  RT_SCOPE_LINK LinkScope = 253
  RT_SCOPE_HOST LinkScope = 254
  RT_SCOPE_NOWHERE LinkScope = 255
)



const (
  IFA_UNSPEC netlink.AttributeType = iota
  IFA_ADDRESS
  IFA_LOCAL
  IFA_LABEL
  IFA_BROADCAST
  IFA_ANYCAST
  IFA_CACHEINFO
  IFA_MULTICAST
  IFA_MAX
)
