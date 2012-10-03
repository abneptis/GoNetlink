package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "syscall"

// A netlink.Socket implements the lowest level of netlink communications.
type Socket struct {
	fd int
}

// Dials a netlink socket.  Usually you do not need permissions for this,
// though specific commands may be rejected.
func Dial(nlf NetlinkFamily) (rwc *Socket, err error) {
	fdno, err := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, int(nlf))
//	err = toErr(errno)
	if err == nil {
		rwc = &Socket{fd: fdno}
	}
	return
}

// Close the netlink socket
func (self *Socket) Close() (err error) {
	syscall.Close(self.fd)
	//err = toErr(errno)
	return
}

// Writes to the netlink socket.  Data should be (1 or more) complete
// netlink frames, as netlink is not friendly w/ fragmentation.
func (self *Socket) Write(in []byte) (n int, err error) {
	n, err = syscall.Write(self.fd, in)
	return
}

// Reads from a netlink socket.  Generally should be a bufio with
// at least an 8k buffer.  More for machines with large routing tables.
func (self *Socket) Read(in []byte) (n int, err error) {
	n, err = syscall.Read(self.fd, in)
	return
}

// Bind the netlink socket to receive multicast messages
func (self *Socket) Bind(pid, groups uint32) (err error) {
	addr := &syscall.SockaddrNetlink{Pid: pid, Groups: groups}
	err = syscall.Bind(self.fd, addr)
	return
}
