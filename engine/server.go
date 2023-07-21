package engine

import (
	"encoding/json"
	"fmt"
	"net"
)

type Server struct {
	Rooms []*Room
	Name  string
	Desc  string
	Port  string
}

func NewServer(name string, desc string, port string) *Server {
	return &Server{
		Name: name,
		Desc: desc,
		Port: port,
	}
}

func (s *Server) Start(port string) error {
	// listen for connections
	l, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		return err
	}

	// close listener when app closes
	defer l.Close()

	fmt.Printf("Server %v started on port %v\n", s.Name, s.Port)

	for {
		// listen for incoming connection
		conn, err := l.Accept()
		if err != nil {
			return err
		}

		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
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

		// handle msg
		var m []byte
		t := data["type"]
		if t == "" {
			m, err = NewMessageBytes("server", "invalid type")
			if err != nil {
				fmt.Println("error creating message", err)
			}
		}

		conn.Write(m)
	}
}
