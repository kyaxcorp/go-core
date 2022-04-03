package base64

import b64 "encoding/base64"

func EncodeString(data string) string {
	return b64.StdEncoding.EncodeToString([]byte(data))
}
