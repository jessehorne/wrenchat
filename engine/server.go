package engine

import (
	"encoding/json"
	"fmt"
	"net"
)

type Server struct {
	Rooms map[string]*Room
	Name  string
	Desc  string
	Port  string
}

var Serv *Server

func NewServer(name string, desc string, port string) *Server {
	return &Server{
		Name: name,
		Desc: desc,
		Port: port,
	}
}

func (s *Server) Start(port string) error {
	Serv = s
	s.Rooms = map[string]*Room{}

	// listen for connections
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return err
	}

	// close listener when app closes
	defer l.Close()

	fmt.Printf("Server %v started on port %v\n", s.Name, s.Port)

	// Add commands
	InitCommands()
	AddCommand("ping", CmdPing)
	AddCommand("msg", CmdSendMsgToRoom)
	AddCommand("create room", CmdCreateRoom)
	AddCommand("join room", CmdJoinRoom)

	userCount := 0
	for {
		// listen for incoming connection
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		userCount += 1
		go handleRequest(userCount, conn)
	}
}

func handleRequest(userID int, conn net.Conn) {
	for {
		buf := make([]byte, 1024)

		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading:", err.Error())
			return
		}

		// parse data received into request
		var data map[string]string
		err = json.Unmarshal(buf[:len], &data)
		if err != nil {
			fmt.Println("error unmarshalling:", err.Error())
			return
		}

		// handle request
		RunCommand(userID, data["cmd"], data, conn)
	}
}
