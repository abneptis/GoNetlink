package main

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "netlink"
import "netlink/rtmanip"
import "netlink/rtnetlink/link"
import "flag"
import "log"

var doUp = flag.Bool("up", false, "Turn interface up")
var doDown = flag.Bool("down", false, "Turn interface up")
var ifName = flag.String("ifname", "", "Interface to use")

func main() {
	flag.Parse()
	nlsock, err := netlink.Dial(netlink.NETLINK_ROUTE)
	if err != nil {
		log.Fatalf("Couldn't dial netlink: %v", err)
	}
	h := netlink.NewHandler(nlsock)
	ec := make(chan error)
	go func() {
		for e := range ec {
			log.Printf("Netlink error: %v", e)
		}
	}()
	go h.Start(ec)
	lf := rtmanip.NewLinkFinder(h)
	l, err := lf.GetLinkByName(*ifName)
	if err == nil {
		if *doDown {
			err = l.SetLinkState(^link.IFF_UP)
			if err != nil {
				log.Printf("Couldn't turn down interface: %v", err)
			}
			l.Refresh()
		}
		if *doUp {
			err = l.SetLinkState(link.IFF_UP)
			if err != nil {
				log.Printf("Couldn't turn up interface: %v", err)
			}
			l.Refresh()
		}
		log.Printf("Link Index: %d", l.LinkIndex())
		log.Printf("Link Name: %s", l.LinkName())
		log.Printf("Link Flags: %s", l.LinkFlags())
		log.Printf("Link MTU: %d", l.LinkMTU())
		log.Printf("Link (l2) Address: %x", l.LinkAddress())
		log.Printf("Link (l2) Broadcast: %x", l.LinkBroadcastAddress())
	} else {
		log.Fatalf("Couldn't get link: %s, %v", *ifName, err)
	}
}
