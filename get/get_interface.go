package get

import (
	pm "agent/report/mess"
)

type ErrNotImplemented struct {
	OS string
}

func (e ErrNotImplemented) Error() string {
	return "not implemented on " + e.OS
}

func IsNotImplemented(err error) bool {
	switch err.(type) {
	case ErrNotImplemented, *ErrNotImplemented:
		return true
	default:
		return false
	}
}

type GetMess interface {
	GetPort() []*pm.Port
	GetIp() []string
	GetOs() []string
	GetHostName() string
}

type Port struct{}
type Ip struct{}
type OsMessage struct{}
