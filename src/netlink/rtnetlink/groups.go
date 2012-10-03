package rtnetlink

const (
	GRP_LINK          = 1 << iota
	GRP_NOTIFY        = 1 << iota
	GRP_NEIGH         = 1 << iota
	GRP_TC            = 1 << iota
	GRP_IPV4_IFADDR   = 1 << iota
	GRP_IPV4_MROUTE   = 1 << iota
	GRP_IPV4_ROUTE    = 1 << iota
	GRP_IPV4_RULE     = 1 << iota
	GRP_IPV6_IFADDR   = 1 << iota
	GRP_IPV6_MROUTE   = 1 << iota
	GRP_IPV6_ROUTE    = 1 << iota
	GRP_IPV6_IFINFO   = 1 << iota
	GRP_DECnet_IFADDR = 1 << iota
	GRP_NOP2          = 1 << iota
	GRP_DECnet_ROUTE  = 1 << iota
	GRP_DECnet_RULE   = 1 << iota
	GRP_NOP4          = 1 << iota
	GRP_IPV6_PREFIX   = 1 << iota
	GRP_IPV6_RULE     = 1 << iota
	GRP_ND_USEROPT    = 1 << iota
	GRP_PHONET_IFADDR = 1 << iota
	GRP_PHONET_ROUTE  = 1 << iota
)
