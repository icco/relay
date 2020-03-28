package lib

import (
	"encoding/json"
	"fmt"
	"io"
)

func ReaderToMessage(b io.Reader) (string, error) {
	var data map[string]string
	if err := json.NewDecoder(b).Decode(&data); err != nil {
		return "", fmt.Errorf("decoding json: %w", err)
	}

	msg := ""
	for k, v := range data {
		msg += fmt.Sprintf("%s: %s\n", k, v)
	}

	return msg, nil
}
