package engine

import (
	"github.com/google/uuid"
	"github.com/jessehorne/wrenchat/util"
	"net"
)

type Room struct {
	Name     string
	Password string
	Address  string // kind of like the mailing address of a room that clients use to route messages
	Users    map[net.Conn]*User
}

func NewRoom(name, password string) *Room {
	return &Room{
		Name:     name,
		Password: password,
		Address:  uuid.NewString(),
	}
}

func (r *Room) SendRawToAll(msg []byte) error {
	for _, u := range r.Users {
		u.Connection.Write(msg)
	}

	return nil
}

func (r *Room) SendEncryptedMessageToAll(from string, msg string) error {
	for _, u := range r.Users {
		data, err := util.EncryptWithPubKey([]byte(msg), u.PublicKey)
		if err != nil {
			return err
		}

		u.Connection.Write(NewMessageBytes("msg", from, string(data)))
	}

	return nil
}
