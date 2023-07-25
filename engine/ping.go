package engine

import (
	"net"
)

func CmdPing(data map[string]string, conn net.Conn) {
	conn.Write(NewMessageBytes("pong", "server", "pong"))
}
