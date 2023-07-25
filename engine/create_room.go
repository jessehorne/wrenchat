package engine

import "net"

func CmdCreateRoom(userID int, data map[string]string, conn net.Conn) {
	name := data["room"]
	password := data["password"]

	if name == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid room name"))
		return
	}

	newRoom := NewRoom(name, password)
	Serv.Rooms[newRoom.Address] = newRoom

	conn.Write(NewMessageBytes("room created", "server", newRoom.Address))
}
