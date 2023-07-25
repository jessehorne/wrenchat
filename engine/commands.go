package engine

import "net"

var Commands map[string]func(map[string]string, net.Conn)

func AddCommand(t string, f func(map[string]string, net.Conn)) {
	Commands[t] = f
}

func InitCommands() {
	Commands = map[string]func(map[string]string, net.Conn){}
}

func RunCommand(t string, data map[string]string, conn net.Conn) {
	_, exists := Commands[t]

	if exists {
		Commands[t](data, conn)
		return
	}

	conn.Write(NewMessageBytes("error", "server", "invalid command"))
}
