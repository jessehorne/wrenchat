package engine

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jessehorne/wrenchat/util"
)

type Room struct {
	Name     string
	Password string
	Address  string // kind of like the mailing address of a room that clients use to route messages
	Users    map[int]*User
}

func NewRoom(name, password string) *Room {
	return &Room{
		Name:     name,
		Password: password,
		Address:  uuid.NewString(),
		Users:    map[int]*User{},
	}
}

func (r *Room) SendRawToAll(msg []byte) error {
	for _, u := range r.Users {
		fmt.Println("sending raw to user", u.ID)
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
