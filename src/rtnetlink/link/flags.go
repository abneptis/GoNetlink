package link

import "strings"

type Flags uint32

const (
  IFF_UP          Flags = 1  << iota
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
  IFF_LOWER_UP
  IFF_DORMANT
  IFF_ECHO
  IFF_VOLATILE    Flags = (IFF_LOOPBACK|IFF_POINTOPOINT|IFF_BROADCAST|IFF_ECHO|IFF_MASTER|IFF_SLAVE|IFF_RUNNING|IFF_LOWER_UP|IFF_DORMANT)
  IFF_QUERY Flags = 0xffffffff
)

func (self Flags)Strings()(out []string){
  if IFF_UP & self  == IFF_UP { out = append(out, "IFF_UP") }
  if IFF_BROADCAST & self  == IFF_BROADCAST { out = append(out, "IFF_BROADCAST") }
  if IFF_DEBUG & self  == IFF_DEBUG { out = append(out, "IFF_DEBUG") }
  if IFF_LOOPBACK & self  == IFF_LOOPBACK { out = append(out, "IFF_LOOPBACK") }
  if IFF_POINTOPOINT & self  == IFF_POINTOPOINT { out = append(out, "IFF_POINTOPOINT") }
  if IFF_NOTRAILERS & self  == IFF_NOTRAILERS { out = append(out, "IFF_NOTRAILERS") }
  if IFF_RUNNING & self  == IFF_RUNNING { out = append(out, "IFF_RUNNING") }
  if IFF_NOARP & self  == IFF_NOARP { out = append(out, "IFF_NOARP") }
  if IFF_PROMISC & self  == IFF_PROMISC { out = append(out, "IFF_PROMISC") }
  if IFF_ALLMULTI & self  == IFF_ALLMULTI { out = append(out, "IFF_ALLMULTI") }
  if IFF_MASTER & self  == IFF_MASTER { out = append(out, "IFF_MASTER") }
  if IFF_SLAVE & self  == IFF_SLAVE { out = append(out, "IFF_SLAVE") }
  if IFF_MULTICAST & self  == IFF_MULTICAST { out = append(out, "IFF_MULTICAST") }
  if IFF_PORTSEL & self  == IFF_PORTSEL { out = append(out, "IFF_PORTSEL") }
  if IFF_AUTOMEDIA & self  == IFF_AUTOMEDIA { out = append(out, "IFF_AUTOMEDIA") }
  if IFF_DYNAMIC & self  == IFF_DYNAMIC { out = append(out, "IFF_DYNAMIC") }
  if IFF_LOWER_UP & self  == IFF_LOWER_UP { out = append(out, "IFF_LOWER_UP") }
  if IFF_DORMANT & self  == IFF_DORMANT { out = append(out, "IFF_DORMANT") }
  if IFF_ECHO & self  == IFF_ECHO { out = append(out, "IFF_ECHO") }
  return
}

func (self Flags)String()(string){
  return strings.Join(self.Strings(), "|")
}
