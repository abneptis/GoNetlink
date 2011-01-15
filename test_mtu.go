package main

import "netlink"
import "netlink/route"
import "log"
//import "syscall"

func main(){
  s, err := netlink.Dial(netlink.NETLINK_ROUTE)
  if err != nil {
    log.Exitf("Couldn't dial netlink message: %v", err)
  }
  h := netlink.NewHandler(s)
  go h.Start()
  msgs, err := route.QueryInterfaces(h)
  //log.Printf("msgs: %+v", msgs)
  if err != nil {
    log.Exitf("Couldn't query: %v", err)
  }
  for m := range(msgs) {
    msg := msgs[m]
    for i := range(msg.Attributes){
      attr := msg.Attributes[i]
      switch attr.AttributeType() {
        default:  log.Printf("Unknown attribute: %X [%X]", attr.AttributeType(), attr.Body())
        case route.IFLA_IFNAME: log.Printf("Name: %s", attr.Body())
        case route.IFLA_MTU: log.Printf("MTU: %d", attr.Body())
      }
    }
    log.Printf("[%d : %+v\n\t%+v", m, msgs[m], msgs[m].Attributes)
  }

  s.Close()
  log.Printf("DONE!")
  return
}
