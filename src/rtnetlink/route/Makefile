include $(GOROOT)/src/Make.inc

TARG=netlink/rtnetlink/route

DEPS=../


GOFILES=\
	flags.go\
	origin.go\
	header.go\
	table.go\
	types.go\


include $(GOROOT)/src/Make.pkg


tester: tester.go _obj/nl3.a route/_obj/nl3/route.a
	$(GC) -o $@.6 $<
	$(LD) -L_obj -Lroute/_obj -o $@ $@.6
