package netlink

import "os"
import "bytes"
import "encoding/binary"

type AttributeType uint16

func (self *AttributeType)Unmarshal(in []byte)( err os.Error){
  var mt uint16
  err = Unmarshal(in, &mt)
  if err == nil { *self = AttributeType(mt) }
  return
}


type Attribute interface {
  AttributeType()(AttributeType)
  Body()([]byte)
}

func (self AttributeType)Marshal()([]byte, os.Error){
  return Marshal(uint16(self))
}

type attr struct {
  _type AttributeType
  _body []byte
}

func (self attr)AttributeType()(AttributeType){ return self._type }
func (self attr)Body()([]byte){ return self._body}

func (self *attr)Unmarshal(in []byte)(err os.Error){
  if len(in) < 2 {
    return os.NewError("Cannot be valid attribute")
  }
  err = Unmarshal(in[0:2], &self._type)
  if err == nil {
    self._body = in[2:]
  }
  return
}

func UnmarshalAttribute(in []byte)(out Attribute, err os.Error){
  a := &attr{}
  err = a.Unmarshal(in)
  if err == nil { out = a }
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
      out = append(out, attr{_type: AttributeType(t), _body:in[pos+4:pos + int(l)]})
      pos += int(l)
      if padding > 0 && pos % padding != 0 {
        pos += padding - (pos % padding)
      }
    } else {
      err = os.NewError("Invalid attribute length.")
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
    if padding > 0 && len(body) % padding != 0 {
      buff.Write(make([]byte, padding - (len(body) % padding)))
    }
  }
  out = buff.Bytes()
  return
}

