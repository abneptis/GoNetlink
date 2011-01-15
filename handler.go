package netlink

import "os"
import "fmt"
import "sync"

type Handler struct {
  sock *Socket
  recipients map[uint32]chan NetlinkMessage
  next_seq uint32
  lock sync.Mutex
}

// This is not atomic...
func (self *Handler)Seq()(out uint32){
  self.lock.Lock()
  out = self.next_seq
  self.next_seq++
  self.lock.Unlock()
  return
}

func NewHandler(sock *Socket)(*Handler){
  return &Handler{sock:sock, recipients: map[uint32]chan NetlinkMessage{}, next_seq: 1}
}

func (self *Handler)SendQuery(msg NetlinkMessage, ch chan NetlinkMessage)(err os.Error){
  if msg.Sequence() == 0 {
    msg.SetSequence(self.Seq())
  }
  if msg.Sequence() == 0 {
    return os.NewError("Failed to set sequence number")
  }
  var ob []byte
  ob, err = Marshal(msg)
  if err == nil {
    self.recipients[msg.Sequence()] = ch
    _, err = self.sock.Write(ob)
  }
  return
}

func (self *Handler)Start(echan chan os.Error){
  rdr := NewReader(self.sock)
  for {
    msg, err := rdr.Read();
    if err == nil {
      if self.recipients[msg.Sequence()] == nil {
        echan <- os.NewError(fmt.Sprintf("GoNetlink: No handler found for sequence number: %d", msg.Sequence()))
        continue
      } else {
        switch msg.MessageType() {
          case NLMSG_DONE:
            close(self.recipients[msg.Sequence()])
            self.recipients[msg.Sequence()] = nil, false
          default: self.recipients[msg.Sequence()] <- msg
        }
      }
    } else {
      echan <- err
    }
  }
  return
}

