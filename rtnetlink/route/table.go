package route

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
