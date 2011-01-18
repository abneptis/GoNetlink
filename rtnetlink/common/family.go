package rtnetlink

import "syscall"

type Family uint8

const (
  AF_UNSPEC Family = syscall.AF_UNSPEC
  AF_INET   Family = syscall.AF_INET
  AF_INET6  Family = syscall.AF_INET6
)

