package netlink
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

type NetlinkFamily uint16 //?

// The netlinkFamily is used for dialing a netlink socket.
// A single socket can only be bound to a single family, and thus
// a single Handler.
const (
  NETLINK_ROUTE NetlinkFamily = iota
  NETLINK_UNUSED
  NETLINK_USERSOCK
  NETLINK_FIREWALL
  NETLINK_INET_DIAG
  NETLINK_NFLOG
  NETLINK_XFRM
  NETLINK_SELINUX
  NETLINK_ISCSI
  NETLINK_AUDIT
  NETLINK_FIB_LOOKUP
  NETLINK_CONNECTOR
  NETLINK_NETFILTER
  NETLINK_IP6_FW
  NETLINK_DNRTMSG
  NETLINK_KOBJECT_UEVENT
  NETLINK_GENERIC
  NETLINK_RESERVED16 // WTF is NETLINK_DM ? 
  NETLINK_SCSITRANSPORT
  NETLINK_ECRYPTFS
)
