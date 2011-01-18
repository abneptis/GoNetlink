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
  nlmsg, err := netlink.NewMessage(rtnetlink.RTM_GETROUTE, netlink.NLM_F_DUMP|netlink.NLM_F_REQUEST, rtmsg, 2)
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
  c := make(chan netlink.Message)
  err = h.SendQuery(*nlmsg, c)
  if err != nil {
    log.Exitf("Couldn't write netlink: %v", err)
  }
  for i := range( c) {
    switch i.Header.MessageType() {
      case rtnetlink.RTM_NEWROUTE:
        hdr := &route.Header{}
        msg := rtnetlink.NewMessage(hdr, nil)
        err = msg.UnmarshalNetlink(i.Body, 4)

        if err == nil {
           log.Printf("Route: %v (%d/%d) TOS: %d; (Table: %v; Origin: %v; Scope: %v; Type: %v; Flags: %v",
                      hdr.AddressFamily(), hdr.AddressDestLength(), hdr.AddressSourceLength(),
                      hdr.TOS(), hdr.RoutingTable(), hdr.RouteOrigin(), hdr.AddressScope(),
                      hdr.RouteType(), hdr.Flags())
           for i := range(msg.Attributes){
             log.Printf("Attribute[%d]: %v", i, msg.Attributes[i])
           }
        } else {
          log.Printf("Unmarshal error: %v", err)
        }
      default:
          log.Printf("Unknown type: %v", i)
    }
  }
}
