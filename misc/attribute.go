package netlink

import "fmt"
import "os"
import "bytes"
import "encoding/binary"

type AttributeType uint16

type Attribute struct {
  Type AttributeType
  Body []byte
}

func (self Attribute)MarshalNetlink(pad int)(out []byte, err os.Error){
  l := len(self.Body)
  out = make([]byte, l + 4)
  binary.LittleEndian.PutUint16(out[0:2], uint16(len(self.Body)+4))
  binary.LittleEndian.PutUint16(out[2:4], uint16(self.Type))
  copy(out[4:], self.Body[0:])
  out = PadBytes(out, pad)
  return
}

func UnmarshalAttributes(in []byte, padding int)(out []Attribute, err os.Error){
  pos := 0
  for pos < len(in) {
    l := binary.LittleEndian.Uint16(in[pos:pos+2])
    if int(l) > len(in) - pos {
      err = os.NewError("Can't parse attribute (too long)")
      break
    }
    if l > 4 {
      t := binary.LittleEndian.Uint16(in[pos+2:pos+4])
      out = append(out, Attribute{Type: AttributeType(t), Body:in[pos+4:pos + int(l)]})
      pos = Reposition(pos + int(l), padding)
    } else {
      err = os.NewError(fmt.Sprintf("Invalid Attributeibute (Len: %d):", l))
      break
    }
  }
  return
}

func MarshalAttributes(in []Attribute, padding int)(out []byte, err os.Error){
  for i := range(in){
    var b []byte
    b, err = in[i].MarshalNetlink(padding)
    out = bytes.Join([][]byte{out, b}, []byte{})
  }
  return
}

