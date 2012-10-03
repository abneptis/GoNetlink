package route
/*
  Copyright (c) 2011, Abneptis LLC. All rights reserved.
  Original Author: James D. Nurmi <james@abneptis.com>

  See LICENSE for details
*/

import "fmt"

type Table byte

const (
        RT_TABLE_UNSPEC Table = 0
        RT_TABLE_COMPAT Table  = 252
        RT_TABLE_DEFAULT Table = 253
        RT_TABLE_MAIN    Table = 254
        RT_TABLE_LOCAL   Table =255
        // Cant be used in Go (wrong size)
        //RT_TABLE_MAX=0xFFFFFFFF
)



var TableStrings = map[Table]string {
  RT_TABLE_UNSPEC: "RT_TABLE_UNSPEC",
  RT_TABLE_COMPAT: "RT_TABLE_COMPAT",
  RT_TABLE_DEFAULT: "RT_TABLE_DEFAULT",
  RT_TABLE_MAIN: "RT_TABLE_MAIN",
  RT_TABLE_LOCAL: "RT_TABLE_LOCAL",
}

func (self Table)String()(out string){
  out = TableStrings[self]
  if out == "" {
    out = fmt.Sprintf("RT_TABLE_%d", self)
  }
  return
}
