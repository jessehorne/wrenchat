package engine

import "net"

func CmdSendMsgToRoom(data map[string]string, conn net.Conn) {
	addr := data["addr"]
	if addr == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid addr"))
		return
	}

	r, exists := Serv.Rooms[addr]
	if !exists {
		conn.Write(NewMessageBytes("error", "server", "room doesn't exist"))
		return
	}

	msg := data["msg"]
	if msg == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid msg"))
		return
	}

	// get user
	u := r.Users[conn]

	r.SendEncryptedMessageToAll(u.ID, msg)
}
