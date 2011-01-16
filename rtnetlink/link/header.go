package link

import "os"


const HEADER_LENGTH = 16
type Header [16]byte

func (self Header)Len()(int) { return HEADER_LENGTH }
func (self *Header)UnmarshalNetlink(in []byte, pad int)(err os.Error){
  if len(in) != HEADER_LENGTH {
    err = os.NewError("Wrong length for Header")
  } else {
    copy(self[0:HEADER_LENGTH], in[0:HEADER_LENGTH])
  }
  return
}

func (self Header)MarshalNetlink(pad int)(out []byte, err os.Error){
  out = self[0:HEADER_LENGTH]
  return
}


/*
              struct ifinfomsg {
                  unsigned char  ifi_family;
                  unsigned short ifi_type;
                  int            ifi_index;
                  unsigned int   ifi_flags;
                  unsigned int   ifi_change;
              };
*/
