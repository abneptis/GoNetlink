package netlink

import "bytes"
import "encoding/binary"
import "io"
import "os"
import "fmt"

type Encoder interface {
  Encode(interface{})(os.Error)
}

type Decoder interface {
  Decode(interface{})(os.Error)
}

type Marshaler interface {
  Marshal()([]byte, os.Error)
}

type Unmarshaler interface {
  Unmarshal([]byte)(os.Error)
}

func UnmarshalLittleEndian(in []byte, ino interface{})(err os.Error){
  err = binary.Read(bytes.NewBuffer(in), binary.LittleEndian, ino)
  return
}

func ReadLittleEndian(r io.Reader, ino interface{})(err os.Error){
  err = binary.Read(r, binary.LittleEndian, ino)
  return
}

func MarshalLittleEndian(in interface{})(out []byte, err os.Error){
  buff := bytes.NewBuffer(nil)
  err = binary.Write(buff, binary.LittleEndian, in)
  if err == nil {
    out = buff.Bytes()
  }
  return
}

func Marshal(in interface{})(out []byte, err os.Error){
  switch t := in.(type) {
    case Marshaler:
      return t.Marshal()
    case int16, int32, int64,
         uint16, uint32, uint64,
         float32, float64:
      out, err = MarshalLittleEndian(t)
    case []byte:
      out = t
    case byte:
      out = []byte{t}
    case int, uint:
      err = os.NewError(fmt.Sprintf("%T can only be marshaled with an exact size (%T[8|16|32|64])", t, t))
    case float:
      err = os.NewError(fmt.Sprintf("%T can only be marshaled with an exact size (%T[32|64])", t, t))
    default:
      err = os.NewError(fmt.Sprintf("Don't know how to marshal %T (be more precise?)", t))
  }
  return
}

type simpleEncoder struct {
  w io.Writer
}

func (self *simpleEncoder)Encode(i interface{})(err os.Error){
  bb, err := Marshal(i)
  if err == nil {
    _, err = self.w.Write(bb)
  }
  return
}

func NewEncoder(w io.Writer)(Encoder){
  return &simpleEncoder{w:w}
}

func Unmarshal(inb []byte, in interface{})(err os.Error){
  switch t := in.(type) {
    case Unmarshaler:
      return t.Unmarshal(inb)
    case *int16, *int32, *int64,
         *uint16, *uint32, *uint64,
         *float32, *float64:
      err = UnmarshalLittleEndian(inb, in)
    case *[]byte:
      *t = inb
    case *byte:
      *t = inb[0]
    case int, uint:
      err = os.NewError(fmt.Sprintf("%T can only be unmarshaled with an exact size (%T[8|16|32|64])", t, t))
    case float:
      err = os.NewError(fmt.Sprintf("%T can only be unmarshaled with an exact size (%T[32|64])", t, t))
    default:
      err = os.NewError(fmt.Sprintf("Don't know how to unmarshal %T (be more precise?)", t))

  }
  return
}
