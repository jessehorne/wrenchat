package engine

import (
	"crypto/rsa"
	"github.com/google/uuid"
	"net"
)

type User struct {
	ID         string
	PublicKey  *rsa.PublicKey
	Connection net.Conn
	Ready      bool // determines if the user is ready to send/receive messages
}

func NewUser(conn net.Conn) *User {
	return &User{
		ID:         uuid.NewString(),
		Connection: conn,
	}
}
