package protocol

type Protocol interface {
	Serve() error
	Close() error
}
