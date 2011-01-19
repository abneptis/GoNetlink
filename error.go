package netlink

import "fmt"
import "encoding/binary"
import "os"
import "syscall"

const ERROR_LENGTH = HEADER_LENGTH+4
type Error [ERROR_LENGTH]byte

func (self Error)Code()(int32){
  return int32(binary.LittleEndian.Uint32(self[0:4]))
}

func (self Error)MarshalNetlink(pad int)(out []byte, err os.Error){
  out = PadBytes(self[0:ERROR_LENGTH], pad)
  return
}

func (self *Error)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) < ERROR_LENGTH {
    return os.NewError(fmt.Sprintf("Invalid netlink error length: %d", len(in)))
  }
  copy(self[0:ERROR_LENGTH], in)
  return
}

func (self Error)String()(string){
  return syscall.Errstr(int(-self.Code()))
}
