package addr

import "strings"

type Flags uint8

// Flags other than SECONDARY and PERMENANT are
// considered 'undocumented' by netlink.
const (
  IFA_F_SECONDARY Flags = 1 << iota
  IFA_F_NODAD
  IFA_F_OPTIMISTIC
  IFA_F_DADFAILED
  IFA_F_HOMEADDRESS
  IFA_F_DEPRECATED
  IFA_F_TENTATIVE
  IFA_F_PERMANENT
  IFA_F_TEMPORARY = IFA_F_SECONDARY
)

func (self Flags)Strings()(out []string){
  if self & IFA_F_SECONDARY == IFA_F_SECONDARY { out = append(out, "IFA_F_SECONDARY") }
  if self & IFA_F_NODAD == IFA_F_NODAD { out = append(out, "IFA_F_NODAD") }
  if self & IFA_F_OPTIMISTIC == IFA_F_OPTIMISTIC { out = append(out, "IFA_F_OPTIMISTIC") }
  if self & IFA_F_DADFAILED == IFA_F_DADFAILED { out = append(out, "IFA_F_DADFAILED") }
  if self & IFA_F_HOMEADDRESS == IFA_F_HOMEADDRESS { out = append(out, "IFA_F_HOMEADDRESS") }
  if self & IFA_F_DEPRECATED == IFA_F_DEPRECATED { out = append(out, "IFA_F_DEPRECATED") }
  if self & IFA_F_TENTATIVE == IFA_F_TENTATIVE { out = append(out, "IFA_F_TENTATIVE") }
  if self & IFA_F_SECONDARY == IFA_F_SECONDARY { out = append(out, "IFA_F_SECONDARY") }
  return
}

func (self Flags)String()(string){
  return strings.Join(self.Strings(), ",")
}
