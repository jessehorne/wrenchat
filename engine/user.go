package engine

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"github.com/jessehorne/wrenchat/util"
	"net"
)

type User struct {
	ID         string
	PublicKey  *rsa.PublicKey
	Nick string
	Connection net.Conn
	Ready      bool // determines if the user is ready to send/receive messages
}

func NewUser(conn net.Conn, key string, nick string) (*User, error) {
	parsedPubKey, err := util.ParsePubKey(key)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:         uuid.NewString(),
		Connection: conn,
		PublicKey: parsedPubKey,
		Nick: nick,
	}, nil
}
