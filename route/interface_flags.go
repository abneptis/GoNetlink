package route

import "os"
import "netlink"

type InterfaceFlags uint32


func (self InterfaceFlags)Marshal()([]byte, os.Error){
  return netlink.Marshal(uint32(self))
}

func (self *InterfaceFlags)Unmarshal(in []byte)( err os.Error){
  var mt uint32
  err = netlink.Unmarshal(in, &mt)
  if err == nil { *self = InterfaceFlags(mt) }
  return
}

const (
    IFF_UP InterfaceFlags = 0x01 << iota
    IFF_BROADCAST
    IFF_DEBUG
    IFF_LOOPBACK
    IFF_POINTOPOINT
    IFF_NOTRAILERS
    IFF_RUNNING
    IFF_NOARP
    IFF_PROMISC
    IFF_ALLMULTI
    IFF_MASTER
    IFF_SLAVE
    IFF_MULTICAST
    IFF_PORTSEL
    IFF_AUTOMEDIA
    IFF_DYNAMIC
)
