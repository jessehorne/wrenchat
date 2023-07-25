package engine

import "encoding/json"

func NewMessageBytes(cmd string, from string, msg string) []byte {
	data := map[string]string{
		"cmd":  cmd,
		"from": from,
		"msg":  msg,
	}

	j, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}

	return j
}

func NewMessageToRoomBytes(cmd string, room string, from string, msg string) []byte {
	data := map[string]string{
		"cmd":  cmd,
		"room": room,
		"from": from,
		"msg":  msg,
	}

	j, err := json.Marshal(data)
	if err != nil {
		return []byte{}
	}

	return j
}
