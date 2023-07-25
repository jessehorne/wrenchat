package engine

import (
	"fmt"
	"github.com/jessehorne/wrenchat/util"
	"net"
)

func CmdJoinRoom(userID int, data map[string]string, conn net.Conn) {
	addr := data["addr"]

	if addr == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid server address"))
		return
	}

	// check if room exists
	r, exists := Serv.Rooms[addr]

	if !exists {
		conn.Write(NewMessageBytes("error", "server", "room doesn't exists"))
		return
	}

	publicKey := data["publicKey"]

	if publicKey == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid public key"))
		return
	}

	nick := data["nick"]

	if nick == "" {
		conn.Write(NewMessageBytes("error", "server", "invalid nickname"))
		return
	}

	newUser, err := NewUser(conn, publicKey, nick)
	if err != nil {
		fmt.Println(err)
		conn.Write(NewMessageBytes("error", "server", "could not create user"))
		return
	}

	fmt.Println("User joined with ID ", userID, newUser.ID)

	r.Users[userID] = newUser
	r.SendRawToAll(NewMessageBytes("user joined", "server", util.PubKeyToString(newUser.PublicKey)))
}
