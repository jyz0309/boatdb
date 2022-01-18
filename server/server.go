package server

import (
	"net"
)

type Server struct {
	addr string
	port string
}

func (svr *Server) Accept(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		c := newClient(conn)
		if err := c.connect(); err != nil {
			return err
		}
		go c.handleConn()
	}
}

func (svr *Server) Listener() (net.Listener, error) {
	listener, err := net.Listen("tcp", svr.Addr())
	if err != nil {
		return nil, err
	}
	return listener, nil
}

func (svr Server) Addr() string {
	return svr.addr + ":" + svr.port
}
