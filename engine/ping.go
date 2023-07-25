package engine

import (
	"net"
)

func CmdPing(userID int, data map[string]string, conn net.Conn) {
	conn.Write(NewMessageBytes("pong", "server", "pong"))
}
