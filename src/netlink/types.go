package netlink
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

// The MessageType consists of 5 fixed "Netlink" types that indicate
// message flow.
// All messageTypes above MIN_TYPE should be considered the 'property' of
// the individual netlink socket family.
type MessageType uint16

// include/linux/netlink.h
const (
  NLMSG_UNSPECIFIED MessageType = iota
  NLMSG_NOOP
  NLMSG_ERROR
  NLMSG_DONE
  NLMSG_OVERRUN
  MIN_TYPE = 0x10
)
