package netlink

import "os"
import "bytes"

type CString string

func (self CString)String()(string){
  return string(self)
}

func (self CString)Marshal()(out []byte, err os.Error){
  buff := bytes.NewBufferString(string(self))
  out = append(buff.Bytes(), 0x00)
  return
}

func (self *CString)Unmarshal(in []byte)(err os.Error){
  if len(in) == 0 {
    return os.NewError("Invalid bytestring to unmarshal (empty)")
  }
  if in[len(in) - 1] != 0x00 {
    return os.NewError("Invalid bytestring to unmarshal (missing NULL)")
  }
  *self = CString(in[0:len(in) - 1])
  return
}
