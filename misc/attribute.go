package netlink

import "os"

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
