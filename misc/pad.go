package netlink

func PadBytes(in []byte, pad int)(out []byte){
  if pad > 0 {
    pblk := (len(in) + (pad - 1)) / pad
    fsize := pblk * pad
    if fsize != len(in) {
      out = make([]byte, fsize)
      copy(out, in)
    } else {
      out = in
    }
  }
  return
}

func Reposition(pos int, pad int)(out int){
  if pad > 0 {
    out = pad * ((pos + (pad - 1) ) / pad)
  } else {
    out = pos
  }
  return
}


