package netlink

import "io"
import "os"
//import "log"
import "bufio"

func NewReader(r io.Reader)(*messageReader){
  return &messageReader{r:bufio.NewReader(r)}
}

type messageReader struct {
  r io.Reader
}

func (self *messageReader)Read()(nlm NetlinkMessage, err os.Error){
  var msglen uint32
  err = ReadLittleEndian(self.r, &msglen)
  if err == nil {
    //log.Printf("Expecting message length: %d", msglen)
    buff := make([]byte, msglen - 4)
    var n int
    n, err = self.r.Read(buff)
    if err == nil {
      if uint32(n) != (msglen - 4) {
        //log.Printf("MISMATCH: Read: %d, Exp: %d", n, msglen)
        err = os.NewError("Invalid netlink (length didn't match read)")
      } else {
        msg := &nlmsg{}
        err = msg.Unmarshal(buff)
        if err == nil { nlm = msg }
        //log.Printf("(In) Netlink Buffer: %X", buff)
      }
    }
  }
  return
}
