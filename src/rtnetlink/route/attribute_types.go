package route

import "netlink"

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/


const (
  RTA_UNSPEC      netlink.AttributeType = iota
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

var AttributeTypeStrings = map[netlink.AttributeType]string {
  RTA_UNSPEC: "RTA_UNSPEC",
  RTA_DST: "RTA_DST",
  RTA_SRC: "RTA_SRC",
  RTA_IIF: "RTA_IIF",
  RTA_OIF: "RTA_OIF",
  RTA_GATEWAY: "RTA_GATEWAY",
  RTA_PRIORITY: "RTA_PRIORITY",
  RTA_PREFSRC: "RTA_PREFSRC",
  RTA_METRICS: "RTA_METRICS",
  RTA_MULTIPATH: "RTA_MULTIPATH",
  RTA_PROTOINFO: "RTA_PROTOINFO",
  RTA_FLOW: "RTA_FLOW",
  RTA_CACHEINFO: "RTA_CACHEINFO",
  RTA_SESSION: "RTA_SESSION",
  RTA_MP_ALGO: "RTA_MP_ALGO",
  RTA_TABLE: "RTA_TABLE",
}
