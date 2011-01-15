package route

type Flags uint32
const (
  RTM_F_NOTIFY   Flags = 0x100
  RTM_F_CLONED   Flags = 0x200
  RTM_F_EQUALIZE Flags = 0x400
  RTM_F_PREFIX   Flags = 0x800
)
