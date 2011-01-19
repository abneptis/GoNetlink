package rtmanip

import "netlink/rtnetlink/link"
import "netlink/rtnetlink"
import "netlink"
import "os"
import "log"

type LinkHandler struct {
  h *netlink.Handler
  cache *rtnetlink.Message
}

func (self *LinkHandler)LinkBroadcastAddress()(s []byte){
  attr, err := self.cache.GetAttribute(link.IFLA_BROADCAST)
  if err == nil {
    s = attr.Body
  }
  return
}

func (self *LinkHandler)LinkAddress()(s []byte){
  attr, err := self.cache.GetAttribute(link.IFLA_ADDRESS)
  if err == nil {
    s = attr.Body
  }
  return
}

func (self *LinkHandler)Refresh()(err os.Error){
  if hdr, ok := self.cache.Header.(*link.Header); ok {
    lf := &linkFinder{h: self.h}
    l, err := lf.GetLinkByID(hdr.InterfaceIndex())
    if err == nil {
      *self.cache = *l.cache
    }
  }
  return
}

func (self *LinkHandler)LinkMTU()(i uint32){
  i, _ = netlink.GetAttributeUint32(self.cache, link.IFLA_MTU)
  return
}

func (self *LinkHandler)SetLinkState(flag link.Flags)(err os.Error){
  var qry *netlink.Message
  if hdr, ok := self.cache.Header.(*link.Header); ok {
    // While rtnetlink(7) says changes should always be IFF_QUERY, it has some
    // behaviours that are undocumented - like limiting actions on SETLINK's to
    // specific FLAGs.
    hdr = link.NewHeader(hdr.InterfaceFamily(), hdr.InterfaceType(), hdr.InterfaceIndex(), flag & link.IFF_UP, link.IFF_UP)
    msg := rtnetlink.Message{Header: hdr}
    qry, err = netlink.NewMessage(rtnetlink.RTM_SETLINK, netlink.NLM_F_ACK|netlink.NLM_F_REQUEST, msg, 4)
  } else {
    err = os.NewError("Cant set link flags (invalid cache)")
  }
  if err != nil { return }
  mch, err := self.h.Query(*qry,1, 4)
  if err == nil {
    for m := range(mch) {
      switch m.Header.MessageType() {
        case netlink.NLMSG_ERROR:
          emsg := &netlink.Error{}
          err = emsg.UnmarshalNetlink(m.Body, 4)
          if err == nil && emsg.Code() != 0 { err = emsg }
        default:
          err = os.NewError("Unexpected netlink message")
          log.Printf("NetlinkError: %v", err)
      }
    }
    close(mch)
  }
  return
}

func (self *LinkHandler)LinkFlags()(flags link.Flags){
  if hdr, ok := self.cache.Header.(*link.Header); ok {
    flags = hdr.Flags()
  }
  return
}

func (self *LinkHandler)LinkIndex()(i uint32){
  if hdr, ok := self.cache.Header.(*link.Header); ok {
    i = hdr.InterfaceIndex()
  }
  return
}

func (self *LinkHandler)LinkName()(s string){
  s, _ = netlink.GetAttributeCString(self.cache, link.IFLA_IFNAME)
  return
}

type LinkFinder interface {
  GetLinkByID(uint32)(*LinkHandler, os.Error)
  GetLinkByName(string)(*LinkHandler, os.Error)
  GetLinks()([]*LinkHandler, os.Error)
}

type linkFinder struct {
  h *netlink.Handler
}

func NewLinkFinder(h *netlink.Handler)(LinkFinder){
  return &linkFinder{h:h}
}

func (self *linkFinder)GetLinkByName(s string)(lh *LinkHandler, err os.Error){
 lhs, err := self.GetLinks()
 for i := range(lhs) {
   if lhs[i].LinkName() == s {
    lh = lhs[i]
    break
   }
 }
 if lh == nil { err = os.NewError("Interface not found")}
 return
}

func (self *linkFinder)GetLinkByID(i uint32)(lh *LinkHandler, err os.Error){
 qry, err := netlink.NewMessage(rtnetlink.RTM_GETLINK, netlink.NLM_F_REQUEST,
             link.NewHeader(rtnetlink.AF_UNSPEC, 0, i, 0,0), 4)
 if err == nil {
   var mch chan netlink.Message
   mch, err = self.h.Query(*qry,1, 4)
   if err == nil {
     for ii := range(mch) {
       switch ii.Header.MessageType(){
         default: err = os.NewError("Unknown message type in response to RTM_GETLINK")
         case netlink.NLMSG_ERROR:
           emsg := &netlink.Error{}
           err = emsg.UnmarshalNetlink(ii.Body, 4)
           if err == nil && emsg.Code() != 0 { err = emsg }
         case rtnetlink.RTM_NEWLINK:
           lhdr := &link.Header{}
           msg := &rtnetlink.Message{Header: lhdr}
           err = msg.UnmarshalNetlink(ii.Body, 4)
           if err == nil { lh = &LinkHandler{h: self.h, cache: msg} }
       }
     }
     close(mch)
   }
 }
 return
}

func (self *linkFinder)GetLinks()(lhs []*LinkHandler, err os.Error){
 qry, err := netlink.NewMessage(rtnetlink.RTM_GETLINK, netlink.NLM_F_REQUEST|netlink.NLM_F_ROOT,
             link.NewHeader(rtnetlink.AF_UNSPEC, 0, 0, 0,0), 4)
 if err != nil { return }
 var mch chan netlink.Message
 mch, err = self.h.Query(*qry,1, 4)
 if err == nil {
   for ii := range(mch) {
     if ii.Header.MessageType() == netlink.NLMSG_DONE { break }
     switch ii.Header.MessageType(){
       default:
         err = os.NewError("Unknown message type in response to RTM_GETLINK")
       case netlink.NLMSG_ERROR:
           emsg := &netlink.Error{}
           err = emsg.UnmarshalNetlink(ii.Body, 4)
           if err == nil && emsg.Code() != 0 { err = emsg }
       case rtnetlink.RTM_NEWLINK:
         lhdr := &link.Header{}
         msg := &rtnetlink.Message{Header: lhdr}
         err = msg.UnmarshalNetlink(ii.Body, 4)
         if err == nil {
           lhs = append(lhs, &LinkHandler{h: self.h, cache: msg})
         }
     }
     if err != nil { log.Printf("Internal netlink failure: %v", err) }
   }
 }
 close(mch)
 return
}
