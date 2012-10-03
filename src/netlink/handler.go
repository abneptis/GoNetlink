package netlink

/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

//import "log"
import (
	"bufio"
	"errors"
)

import "fmt"
import "sync"

// A handler implements a simple Mux for netlink messages, ensuring
// each query gets a unique sequence number and a channel to collect responses.
type Handler struct {
	sock       *Socket
	recipients map[uint32]chan Message
	next_seq   uint32
	lock       sync.Mutex
}

// Used as an atomic counter for sequence numbering.
// No check is made to see that sequences aren't still in use on roll-over.
func (self *Handler) Seq() (out uint32) {
	self.lock.Lock()
	out = self.next_seq
	self.next_seq++
	self.lock.Unlock()
	return
}

func NewHandler(sock *Socket) *Handler {
	return &Handler{sock: sock, recipients: map[uint32]chan Message{}, next_seq: 1}
}

// Send a message.  If SequenceNumber is unset, Seq() will be used
// to generate one.
func (self *Handler) Query(msg Message, l int, pad int) (ch chan Message, err error) {
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

// Usually called in a goroutine, Start spawns a thread
// that demux's incoming netlink responses.
// Echan is used to report internal netlink errors, and may
// be set to nil (but you will likely miss bugs!)
func (self *Handler) Start(echan chan error) {
	r := bufio.NewReader(self.sock)
	for {
		msg, err := ReadMessage(r, 4)
		if err == nil {
			if self.recipients[msg.Header.MessageSequence()] == nil {
				if nil != echan {
					echan <- errors.New(fmt.Sprintf("GoNetlink: No handler found for seq %d",
						msg.Header.MessageSequence()))
				}
				continue
			} else {
				self.recipients[msg.Header.MessageSequence()] <- *msg
				if msg.Header.MessageFlags()&NLM_F_MULTI == 0 {
					close(self.recipients[msg.Header.MessageSequence()])
					delete(self.recipients, msg.Header.MessageSequence())
				}
			}
		} else {
			if nil != echan {
				echan <- err
			}
		}
	}
	return
}
