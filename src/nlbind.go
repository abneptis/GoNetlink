package netlink

/*
#include <sys/socket.h>
#include <linux/netlink.h>
#include <string.h>

int nlbind(int fd, unsigned int pid, unsigned int groups) {
	struct sockaddr_nl sa;
	memset(&sa, 0, sizeof(sa));

	sa.nl_family = AF_NETLINK;
	sa.nl_pid = pid;
	sa.nl_groups = groups;

	return bind(fd,(struct sockaddr *) &sa, sizeof(sa));
}
*/
import "C"

import "os"

// Bind the socket to the specified netlink multicast groups
func nlBind(sock *Socket, pid, groups uint32) (err os.Error) {
	_, err = C.nlbind(C.int(sock.fd), C.uint(pid), C.uint(groups))
	return
}
