package netlink

import "os"
import "syscall"

// A netlink.Socket implements the lowest level of netlink communications.
type Socket struct {
  fd int
}

func toErr(eno int)(err os.Error){
  if eno != 0 { err = os.NewError(syscall.Errstr(eno))}
  return
}

// Dials a netlink socket.  Usually you do not need permissions for this,
// though specific commands may be rejected.
func Dial(nlf NetlinkFamily)(rwc *Socket, err os.Error){
  fdno, errno := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, int(nlf))
  err = toErr(errno)
  if err == nil {
    rwc = &Socket{fd:fdno}
  }
  return
}

// Close the netlink socket
func (self *Socket)Close()(err os.Error){
  errno := syscall.Close(self.fd)
  err = toErr(errno)
  return
}

// Writes to the netlink socket.  Data should be (1 or more) complete
// netlink frames, as netlink is not friendly w/ fragmentation.
func (self *Socket)Write(in []byte)(n int, err os.Error){
  n, errno := syscall.Write(self.fd, in)
  err = toErr(errno)
  return
}

// Reads from a netlink socket.  Generally should be a bufio with
// at least an 8k buffer.  More for machines with large routing tables.
func (self *Socket)Read(in []byte)(n int, err os.Error){
  n, errno := syscall.Read(self.fd, in)
  err = toErr(errno)
  return
}

