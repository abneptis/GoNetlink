package netlink
//import "log"
import "bufio"
import "os"
import "fmt"
import "sync"

type Handler struct {
  sock *Socket
  recipients map[uint32]chan Message
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
  return &Handler{sock:sock, recipients: map[uint32]chan Message{}, next_seq: 1}
}

func (self *Handler)Query(msg Message, l int, pad int)(ch chan Message, err os.Error){
  if msg.Header.MessageSequence() == 0 {
    msg.Header.SetMessageSequence(self.Seq())
  }
  ob, err := msg.MarshalNetlink(pad)
  if err == nil {
    ch = make(chan Message, l)
    self.recipients[msg.Header.MessageSequence()] = ch
    _, err = self.sock.Write(ob)
  }
  return
}

func (self *Handler)Start(echan chan os.Error){
  r := bufio.NewReader(self.sock)
  for {
    msg, err := ReadMessage(r, 4)
    if err == nil {
      if self.recipients[msg.Header.MessageSequence()] == nil {
        echan <- os.NewError(fmt.Sprintf("GoNetlink: No handler found for sequence number: %d", msg.Header.MessageSequence()))
        continue
      } else {
        self.recipients[msg.Header.MessageSequence()] <- *msg
        if msg.Header.MessageFlags() & NLM_F_MULTI == 0 {
            close(self.recipients[msg.Header.MessageSequence()])
            self.recipients[msg.Header.MessageSequence()] = nil, false
        }
      }
    } else {
      echan <- err
    }
  }
  return
}

