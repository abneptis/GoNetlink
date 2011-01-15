package main

import "netlink"
import "log"

func main(){
  i := 0x01234567
  out, err := netlink.Marshal(int8(i))
  if err == nil {
    log.Printf("Marshal of %X => [%X]", i, out)
  } else {
    log.Exitf("Couldn't marshal: %v", err)
  }
}
