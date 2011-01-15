package route

type Scope byte
const (
  RT_SCOPE_UNIVERSE Scope = 00
  RT_SCOPE_SITE     Scope = 200
  RT_SCOPE_LINK     Scope = 253
  RT_SCOPE_HOST     Scope = 254
  RT_SCOPE_NOWHERE  Scope = 255
)
