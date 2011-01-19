package rtnetlink

import "os"

// An RTNetlink Header describes the structures
// between the netlink header and the nlattributes.
type Header interface {
  Len()(int) // The (unpadded) length of the Header.
  UnmarshalNetlink([]byte, int)(os.Error)
  MarshalNetlink(int)([]byte, os.Error)
}
