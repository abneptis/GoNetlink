package netlink

// Returns a padded version of bytes
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

// Returns where the position should be to 
// read a new object.
func Reposition(pos int, pad int)(out int){
  if pad > 0 {
    out = pad * ((pos + (pad - 1) ) / pad)
  } else {
    out = pos
  }
  return
}


