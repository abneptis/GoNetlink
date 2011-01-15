package rtnetlink
// netlink/route

import "netlink"
import "os"
import "net"

type LinkAddress net.IP
type LocalLinkAddress LinkAddress
func (self LocalLinkAddress)AttributeType()(netlink.AttributeType){ return IFA_LOCAL }
func (self LocalLinkAddress)Body()(out []byte){
  out, _ = self.Marshal()
  return
}
func (self LocalLinkAddress)Marshal()(out []byte, err os.Error){
  return netlink.Marshal([]byte(self))
}

func (self LinkAddress)AttributeType()(netlink.AttributeType){ return IFA_ADDRESS }
func (self LinkAddress)Marshal()(out []byte, err os.Error){
  return netlink.Marshal([]byte(self))
}

func (self LinkAddress)Body()(out []byte){
  out, _ = self.Marshal()
  return
}

