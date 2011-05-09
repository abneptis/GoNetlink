package rtnetlink

const (
	GRP_NONE          = iota
	GRP_LINK          = iota
	GRP_NOTIFY        = iota
	GRP_NEIGH         = iota
	GRP_TC            = iota
	GRP_IPV4_IFADDR   = iota
	GRP_IPV4_MROUTE   = iota
	GRP_IPV4_ROUTE    = iota
	GRP_IPV4_RULE     = iota
	GRP_IPV6_IFADDR   = iota
	GRP_IPV6_MROUTE   = iota
	GRP_IPV6_ROUTE    = iota
	GRP_IPV6_IFINFO   = iota
	GRP_DECnet_IFADDR = iota
	GRP_NOP2          = iota
	GRP_DECnet_ROUTE  = iota
	GRP_DECnet_RULE   = iota
	GRP_NOP4          = iota
	GRP_IPV6_PREFIX   = iota
	GRP_IPV6_RULE     = iota
	GRP_ND_USEROPT    = iota
	GRP_PHONET_IFADDR = iota
	GRP_PHONET_ROUTE  = iota
)
