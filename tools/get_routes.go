package main

/* The output of this utility is JSON, but not intended to be easily human-readable.
   At this moment it exists to test the rtnetlink/route subsystem */

import "os"
import "netlink/rtnetlink/route"
import "netlink/rtnetlink"
import "log"
import "netlink"

func main(){
  rtmsg := route.NewMessage(0,0,0,0,0,0,0,0,0)
  nlmsg, err := netlink.NewMessage2(rtnetlink.RTM_GETROUTE, netlink.NLM_F_DUMP|netlink.NLM_F_REQUEST, rtmsg)
  if err != nil {
    log.Exitf("Couldn't construct message: %v", err)
  }
  nlsock, err := netlink.Dial(netlink.NETLINK_ROUTE)
  if err != nil {
    log.Exitf("Couldn't dial netlink: %v", err)
  }
  h := netlink.NewHandler(nlsock)
  ec := make(chan os.Error)
  go h.Start(ec)
  c := make(chan netlink.NetlinkMessage)
  err = h.SendQuery(nlmsg, c)
  if err != nil {
    log.Exitf("Couldn't write netlink: %v", err)
  }
  for i := range( c) {
    switch i.MessageType() {
      case rtnetlink.RTM_NEWROUTE:
        msg := &route.RoutingMessage{}
        err = msg.UnmarshalNetlink(i.Body(), 4)
        if err == nil {
          log.Printf("NLMsg: %v", msg)
        } else {
          log.Printf("Unmarshal error: %v", err)
        }
      default:
          log.Printf("Unknown type: %v", i)
    }
  }
}
