package ads

import (
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func decodeADS(message []byte) (string, error) {
	// Decode from windows-1252
	decoder := charmap.Windows1252.NewDecoder()
	decoded, err := decoder.String(string(message))
	if err != nil {
		return "", err
	}

	// Strip whitespace and null terminators
	return strings.Trim(decoded, " \t\n\r\x00"), nil
}
