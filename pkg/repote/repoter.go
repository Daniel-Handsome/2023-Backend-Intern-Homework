package repote

type Reporter interface {
	Send(msg string) error
}
