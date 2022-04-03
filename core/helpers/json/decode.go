package json

import "encoding/json"

func Decode(v string, to interface{}) error {
	if _err := json.Unmarshal([]byte(v), &to); _err != nil {
		return _err
	}
	return nil
}
