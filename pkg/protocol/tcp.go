package protocol

import (
	"fmt"
	"net"
)

func NewTcpProtocol(port int, handler func(net.Conn) error) Protocol {
	return &tcpProtocol{
		port:    port,
		handler: handler,
	}
}

type tcpProtocol struct {
	port     int
	listener net.Listener
	handler  func(net.Conn) error
}

func (t *tcpProtocol) Serve() error {
	address := fmt.Sprintf(":%d", t.port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	t.listener = listen
	for {
		conn, err := t.listener.Accept()
		if err != nil {
			return err
		}
		err = t.handler(conn)
		if err != nil {
			return err
		}
	}
}

func (t *tcpProtocol) Close() error {
	if t.listener != nil {
		return t.listener.Close()
	}
	return nil
}
