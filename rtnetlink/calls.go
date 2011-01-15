package rtnetlink

import "syscall"
import "netlink"
import "net"
import "os"
import "fmt"

func Dial()(*netlink.Socket, os.Error){
  return netlink.Dial(netlink.NETLINK_ROUTE)
}

func FindInterfaceByName(h *netlink.Handler, n string)(ifmsg InterfaceInformationMessage, err os.Error){
  var found = false
  ifs, err := QueryInterfaces(h)
  if err == nil {
    for i := range(ifs){
      var attr netlink.Attribute
      attr, err = ifs[i].GetAttribute(IFLA_IFNAME)
      // Does such a beast as an unnamed interface exist?
      if err != nil { break }
      if string(attr.Body()) == (n + "\x00") {
        ifmsg, found = ifs[i], true
        break
      }
    }
  }
  if !found && err == nil { err = os.NewError("Can't find interface") }
  return
}


func UpdateInterface(h *netlink.Handler, rtmsg InterfaceInformationMessage)(err os.Error){
  nlmsg, err := netlink.NewMessage(RTM_NEWLINK, netlink.NLM_F_ACK| netlink.NLM_F_REQUEST , rtmsg)
  if err != nil { return }

  c := make(chan netlink.NetlinkMessage)
  err = h.SendQuery(nlmsg, c)

  if err != nil {
   close(c)
   return
  }

  err = h.SendQuery(nlmsg, c)
  for msg := range(c) {
    if msg.MessageType() == netlink.NLMSG_ERROR {
      close(c)
      var code int32
      err = netlink.Unmarshal(msg.Body()[0:4], &code)
      if err == nil && code != 0 {
        err = os.NewError(fmt.Sprintf("Upper Netlink error: [%X]", code ))
      }
      break
    }
    err = os.NewError("Unexpected netlink message type")
    break
  }
  return
}

func QueryInterfaces(h *netlink.Handler)(out []InterfaceInformationMessage, err os.Error){
  rtmsg := NewInterfaceInfoMesage(0, 0, 0, 0, nil )// []netlink.Attribute{route.MTU(1480)})
  nlmsg, err := netlink.NewMessage(RTM_GETLINK, netlink.NLM_F_ECHO | netlink.NLM_F_REQUEST | netlink.NLM_F_ROOT, rtmsg)
  if err != nil { return }

  c := make(chan netlink.NetlinkMessage)
  err = h.SendQuery(nlmsg, c)

  if err != nil {
   close(c)
   return
  }

  for msg := range(c) {
    switch msg.MessageType() {
      case RTM_NEWLINK:
        rtm_msg := &InterfaceInformationMessage{}
        err = rtm_msg.Unmarshal(msg.Body())
        if err == nil { out = append(out, *rtm_msg) }
      case netlink.NLMSG_ERROR:
        close(c)
        err = os.NewError(fmt.Sprintf("Upper Netlink error: [%+v]", msg ))
      default:
        err = os.NewError("Unexpected netlink message type")
    }
    if err != nil { break }
  }
  return
}

func QueryInterfaceAddresses(h *netlink.Handler)(out []InterfaceAddressMessage, err os.Error){
  rtmsg := NewInterfaceAddrMesage(0, 0, 0, 0, 0, nil )// []netlink.Attribute{route.MTU(1480)})
  nlmsg, err := netlink.NewMessage(RTM_GETADDR, netlink.NLM_F_ECHO | netlink.NLM_F_REQUEST | netlink.NLM_F_ROOT, rtmsg)
  if err != nil { return }

  c := make(chan netlink.NetlinkMessage)
  err = h.SendQuery(nlmsg, c)

  if err != nil {
   close(c)
   return
  }

  for msg := range(c) {
    switch msg.MessageType() {
      case RTM_NEWADDR:
        rtm_msg := &InterfaceAddressMessage{}
        err = rtm_msg.Unmarshal(msg.Body())
        if err == nil { out = append(out, *rtm_msg) }
      case netlink.NLMSG_ERROR:
        close(c)
        err = os.NewError(fmt.Sprintf("Upper Netlink error: [%+v]", msg ))
      default:
        err = os.NewError("Unexpected netlink message type")
    }
    if err != nil { break }
  }
  return
}

func AddInterfaceAddress(h *netlink.Handler, ifname string, addr net.IP, mask int)(err os.Error){
  var fam, prefix byte
  if addr.To4() != nil {
    if mask < 0 || mask > 32 {
      return os.NewError("Invalid IPv4 Mask")
    }
    fam = syscall.AF_INET
    prefix = byte(mask)
    addr = addr.To4()
  } else if addr.To16() != nil {
    if mask < 0 || mask > 128 {
      return os.NewError("Invalid IPv6 Mask")
    }
    fam = syscall.AF_INET6
    prefix = byte(mask)
  }
  _if, err := FindInterfaceByName(h, ifname)
  if err != nil { return }
  rtmsg := NewInterfaceAddrMesage(fam, prefix, 0, RT_SCOPE_HOST, _if.Index, []netlink.Attribute{LinkAddress(addr), LocalLinkAddress(addr)})
  nlmsg, err := netlink.NewMessage(RTM_NEWADDR,netlink.NLM_F_REQUEST | netlink.NLM_F_EXCL | netlink.NLM_F_ATOMIC | netlink.NLM_F_ACK, rtmsg)
  if err != nil { return }

  c := make(chan netlink.NetlinkMessage)
  err = h.SendQuery(nlmsg, c)

  if err != nil {
   close(c)
   return
  }

  for msg := range(c) {
    if msg.MessageType() == netlink.NLMSG_ERROR {
      close(c)
      var code int32
      err = netlink.Unmarshal(msg.Body()[0:4], &code)
      if err == nil && code != 0 {
        err = os.NewError(fmt.Sprintf("Upper Netlink error: [%X]", code ))
      }
      break
    }
    err = os.NewError("Unexpected netlink message type")
    break
  }
  return
}

