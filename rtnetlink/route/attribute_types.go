package route

type AttributeType uint16

const (
  RTA_UNSPEC      AttribyteType = iota
  RTA_DST
  RTA_SRC
  RTA_IIF
  RTA_OIF
  RTA_GATEWAY
  RTA_PRIORITY
  RTA_PREFSRC
  RTA_METRICS
  RTA_MULTIPATH
  RTA_PROTOINFO // Deprecated (?)
  RTA_FLOW
  RTA_CACHEINFO
  RTA_SESSION // Deprecated (?)
  RTA_MP_ALGO // Deprecated (?)
  RTA_TABLE
)
