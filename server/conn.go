package server

import "net"

type client struct {
	conn net.Conn
}

func newClient(conn net.Conn) *client {
	return &client{conn: conn}
}

// connect with mysql client, include handshake and auth.
func (c *client) connect() error {
	return nil
}

func (c *client) handleConn() {

}
