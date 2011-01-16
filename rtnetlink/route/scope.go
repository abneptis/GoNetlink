package route

import "fmt"

type Scope byte
const (
  RT_SCOPE_UNIVERSE Scope = 00
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

func (self Scope)String()(out string){
  out = ScopeStrings[self]
  if out == "" {
    out = fmt.Sprintf("RT_SCOPE_%d", self)
  }
  return
}
