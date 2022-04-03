package base64

import b64 "encoding/base64"

func DecodeString(data string) string {
	decoded, _ := b64.StdEncoding.DecodeString(data)
	return string(decoded)
}
