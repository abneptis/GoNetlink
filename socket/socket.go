package netlink

import "os"
import "syscall"

type Socket struct {
  fd int
}

func toErr(eno int)(err os.Error){
  if eno != 0 { err = os.NewError(syscall.Errstr(eno))}
  return
}

func Dial(nlf NetlinkFamily)(rwc *Socket, err os.Error){
  fdno, errno := syscall.Socket(syscall.AF_NETLINK, syscall.SOCK_DGRAM, int(nlf))
  err = toErr(errno)
  if err == nil {
    rwc = &Socket{fd:fdno}
  }
  return
}

func (self *Socket)Close()(err os.Error){
  errno := syscall.Close(self.fd)
  err = toErr(errno)
  return
}


func (self *Socket)Write(in []byte)(n int, err os.Error){
  n, errno := syscall.Write(self.fd, in)
  err = toErr(errno)
  return
}

func (self *Socket)Read(in []byte)(n int, err os.Error){
  n, errno := syscall.Read(self.fd, in)
  err = toErr(errno)
  return
}

