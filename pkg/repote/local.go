package repote

import (
	"runtime/debug"

	log "github.com/sirupsen/logrus"
)

type Local struct {
	Err string
}

func NewLocal() *Local {
	return &Local{}
}

func (s *Local) Send(msg string) error {
	log.WithField("error message", msg).
		WithField("Stack", string(debug.Stack())).
		Error("grpc api panic")

	return nil
}
