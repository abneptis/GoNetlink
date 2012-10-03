include $(GOROOT)/src/Make.inc

TARG=netlink/rtnetlink/addr

DEPS=../


GOFILES=\
	flags.go\
	header.go\
	attribute_types.go\


include $(GOROOT)/src/Make.pkg


tester: tester.go _obj/nl3.a route/_obj/nl3/route.a
	$(GC) -o $@.6 $<
	$(LD) -L_obj -Lroute/_obj -o $@ $@.6
