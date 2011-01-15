include $(GOROOT)/src/Make.inc

TARG=netlink/rtnetlink

DEPS=../


GOFILES=\
	calls.go\
	interface/interface_info.go\
	interface/interface_addr.go\
	interface/interface_flags.go\
	link/link_attributes.go\
	link/link_address.go\
	message_types.go\
	mtu.go\


include $(GOROOT)/src/Make.pkg


tester: tester.go _obj/nl3.a route/_obj/nl3/route.a
	$(GC) -o $@.6 $<
	$(LD) -L_obj -Lroute/_obj -o $@ $@.6
