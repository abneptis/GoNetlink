package rtnetlink

import "os"

type Header interface {
  Len()(int)
  UnmarshalNetlink([]byte, int)(os.Error)
  MarshalNetlink(int)([]byte, os.Error)
}
