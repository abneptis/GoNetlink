package rtnetlink
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "fmt"

// The scope of an address/route
type Scope byte

const (
  RT_SCOPE_UNIVERSE Scope = 0
  RT_SCOPE_SITE     Scope = 200
  RT_SCOPE_LINK     Scope = 253
  RT_SCOPE_HOST     Scope = 254
  RT_SCOPE_NOWHERE  Scope = 255
)


var ScopeStrings = map[Scope]string {
  RT_SCOPE_UNIVERSE: "RT_SCOPE_UNIVERSE",
  RT_SCOPE_SITE: "RT_SCOPE_SITE",
  RT_SCOPE_LINK: "RT_SCOPE_LINK",
  RT_SCOPE_HOST: "RT_SCOPE_HOST",
  RT_SCOPE_NOWHERE: "RT_SCOPE_NOWHERE",
}

// Will return a string identifying the Scope,
// if a 'user' table, RT_SCOPE_%d will be used.
func (self Scope)String()(out string){
  out = ScopeStrings[self]
  if out == "" {
    out = fmt.Sprintf("RT_SCOPE_%d", self)
  }
  return
}
