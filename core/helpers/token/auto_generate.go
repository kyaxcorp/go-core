package token

import "github.com/kyaxcorp/go-core/core/helpers/str"

// AutoGenerated -> will create a random token with prefix auto-generated-xxxxxxxxxxxxxxxxxxxxx
func AutoGenerated(length int) string {
	return "auto-generated-" + str.RandomGenerate(length)
}
