package netlink


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
