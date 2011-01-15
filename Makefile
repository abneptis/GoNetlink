include $(GOROOT)/src/Make.inc

TARG=netlink


GOFILES=\
	misc/cstring.go\
	misc/attribute.go\
	message/flags.go\
	message/types.go\
	message/header.go\
	message/message.go\
	message/message_reader.go\
	socket/family.go\
	socket/socket.go\
	socket/address.go\
	marshal.go\
	handler.go\


include $(GOROOT)/src/Make.pkg


