package engine

import "encoding/json"

func NewMessageBytes(from string, msg string) ([]byte, error) {
	data := map[string]string{
		"from": from,
		"msg":  msg,
	}

	j, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return j, nil
}
