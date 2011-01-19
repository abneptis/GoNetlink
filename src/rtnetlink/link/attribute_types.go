package link
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "netlink"

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

