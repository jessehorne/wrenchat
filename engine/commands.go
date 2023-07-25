package engine

import "net"

var Commands map[string]func(int, map[string]string, net.Conn)

func AddCommand(t string, f func(int, map[string]string, net.Conn)) {
	Commands[t] = f
}

func InitCommands() {
	Commands = map[string]func(int, map[string]string, net.Conn){}
}

func RunCommand(userID int, t string, data map[string]string, conn net.Conn) {
	_, exists := Commands[t]

	if exists {
		Commands[t](userID, data, conn)
		return
	}

	conn.Write(NewMessageBytes("error", "server", "invalid command"))
}
