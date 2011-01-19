include $(GOROOT)/src/Make.inc

all: get_addrs get_links get_routes ifmanip


get_addrs.$(O): get_addrs.go
	$(GC) -o $@ $<

get_links.$(O): get_links.go
	$(GC) -o $@ $<

get_routes.$(O): get_routes.go
	$(GC) -o $@ $<

ifmanip.$(O): ifmanip.go
	$(GC) -o $@ $<

get_addrs: get_addrs.$(O)
	$(LD) -o $@ $<

get_links: get_links.$(O)
	$(LD) -o $@ $<

get_routes: get_routes.$(O)
	$(LD) -o $@ $<

ifmanip: ifmanip.$(O)
	$(LD) -o $@ $<

clean:
	rm -f get_links get_routes get_addrs *.6 *.5 *.8
