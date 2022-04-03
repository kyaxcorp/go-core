package json

import (
	"encoding/json"
	"github.com/KyaXTeam/go-core/v2/core/helpers/conv"
)

func Encode(v interface{}) (string, error) {
	encoded, _err := json.Marshal(v)
	if _err != nil {
		return "", _err
	}
	return conv.BytesToStr(encoded), nil
}

func EncodeBeautify(v interface{}) (string, error) {
	encoded, _err := json.MarshalIndent(v, "", " ")
	if _err != nil {
		return "", _err
	}
	return conv.BytesToStr(encoded), nil
}
