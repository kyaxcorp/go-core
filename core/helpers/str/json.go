package str

import "github.com/KyaXTeam/go-core/v2/core/helpers/json"

func IsJSON(str string) bool {
	return json.IsJSON(str)
}
