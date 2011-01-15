package rtnetlink
// netlink/route

import "netlink"
import "os"

type MTU uint32
func (self MTU)AttributeType()(netlink.AttributeType){ return IFLA_MTU }
func (self MTU)Marshal()([]byte, os.Error){ return netlink.Marshal(uint32(self)) }

func (self MTU)Body()(out []byte){
  out, _ = self.Marshal()
  return
}

