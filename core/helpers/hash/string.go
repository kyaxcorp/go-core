package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

func Sha256(s string) string {
	//return hex.EncodeToString(sha256.Sum256([]byte(s))[:]))
	sum := sha256.Sum256([]byte(s))

	//return string(sum[:])
	return hex.EncodeToString(sum[:])
}

func MD5(s string) string {
	sum := md5.Sum([]byte(s))
	//return string(sum[:])
	return hex.EncodeToString(sum[:])
}

func Sha1(s string) string {
	sum := sha1.Sum([]byte(s))
	//return string(sum[:])
	return hex.EncodeToString(sum[:])
}
