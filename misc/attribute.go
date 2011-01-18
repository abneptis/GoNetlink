package netlink

import "fmt"
import "os"
import "bytes"
import "encoding/binary"

type AttributeType uint16

type Attribute interface {
  AttributeType()(AttributeType)
  Body()([]byte)
}

type attr struct {
  _type AttributeType
  _body []byte
}

func (self attr)AttributeType()(AttributeType){ return self._type }
func (self attr)Body()([]byte){ return self._body}

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
      out = append(out, attr{_type: AttributeType(t), _body:in[pos+4:pos + int(l)]})
      pos = Reposition(pos + int(l), padding)
    } else {
      err = os.NewError(fmt.Sprintf("Invalid attribute (Len: %d):", l))
      break
    }
  }
  return
}

func MarshalAttributes(in []Attribute, padding int)(out []byte, err os.Error){
  buff := bytes.NewBuffer(nil)
  for i := range(in){
    ahdr := [4]byte{}
    body := in[i].Body()
    binary.LittleEndian.PutUint16(ahdr[2:4], uint16(in[i].AttributeType()))
    binary.LittleEndian.PutUint16(ahdr[0:2], uint16(len(body)) + 4)
    buff.Write(ahdr[0:])
    buff.Write(body)
  }
  out = PadBytes(buff.Bytes(), padding)
  return
}

