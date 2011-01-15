package route

type Origin byte
const (
  RTPROT_UNSPEC     Origin = iota
  RTPROT_REDIRECT
  RTPROT_KERNEL
  RTPROT_BOOT
  RTPROT_STATIC
)
