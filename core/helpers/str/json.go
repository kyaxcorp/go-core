package str

import "github.com/kyaxcorp/go-core/core/helpers/json"

func IsJSON(str string) bool {
	return json.IsJSON(str)
}
