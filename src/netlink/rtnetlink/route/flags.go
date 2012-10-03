package route

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "strings"

type Flags uint32
const (
  RTM_F_NOTIFY   Flags = 0x100
  RTM_F_CLONED   Flags = 0x200
  RTM_F_EQUALIZE Flags = 0x400
  RTM_F_PREFIX   Flags = 0x800
)

func (self Flags)Strings()(out []string){
  if RTM_F_NOTIFY & self  == RTM_F_NOTIFY {
    out = append(out, "RTM_F_NOTIFY")
  }
  if RTM_F_CLONED & self  == RTM_F_CLONED {
    out = append(out, "RTM_F_CLONED")
  }
  if RTM_F_EQUALIZE & self  == RTM_F_EQUALIZE {
    out = append(out, "RTM_F_EQUALIZE")
  }
  if RTM_F_PREFIX & self  == RTM_F_PREFIX {
    out = append(out, "RTM_F_PREFIX")
  }
  return
}

func (self Flags)String()(out string){
  return strings.Join(self.Strings(), ",")
}
