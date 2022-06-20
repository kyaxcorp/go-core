package hash

import (
	"github.com/google/uuid"
	"time"
)

func newRandom() string {
	id, _err := uuid.NewRandom()
	if _err == nil {
		return id.String()
	}
	// if failed to get a random uuid, let's return the current time...
	return time.Now().String()
}

func RandSha256() string {
	return Sha256(newRandom())
}

func RandMD5(s string) string {
	return MD5(newRandom())
}

func RandSha1(s string) string {
	return Sha1(newRandom())
}
