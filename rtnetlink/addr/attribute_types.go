package addr

import "netlink"

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

