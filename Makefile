include $(GOROOT)/src/Make.inc

TARG=netlink


GOFILES=\
	error.go\
	misc/attribute.go\
	misc/pad.go\
	misc/attr_helpers.go\
	message/flags.go\
	message/types.go\
	message/header.go\
	message/message.go\
	socket/family.go\
	socket/socket.go\
	socket/address.go\
	handler.go\


include $(GOROOT)/src/Make.pkg

