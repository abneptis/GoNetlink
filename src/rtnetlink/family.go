package rtnetlink

import "syscall"

// Family is already defined in syscall, but we re-define
// them here for type-safety.
type Family uint8

const (
  AF_UNSPEC Family = syscall.AF_UNSPEC
  AF_INET   Family = syscall.AF_INET
  AF_INET6  Family = syscall.AF_INET6
)


// Returns the String representation of the Family,
// or "" if the family name is unknown.
func (self Family)String()(out string){
  switch self {
    default: out = ""
    case AF_UNSPEC: out = "AF_UNSPEC"
    case AF_INET: out = "AF_INET"
    case AF_INET6: out = "AF_INET6"
  }
  return
}
