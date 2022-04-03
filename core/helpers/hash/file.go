package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

func FileMD5() {

}

func FileSha256(path string) (string, error) {
	f, _err := os.Open(path)
	if _err != nil {
		return "", _err
	}
	defer f.Close()

	h := sha256.New()
	if _, _err := io.Copy(h, f); _err != nil {
		// log.Fatal(_err)
		return "", _err
	}

	// fmt.Printf("%x", h.Sum(nil))
	return hex.EncodeToString(h.Sum(nil)), nil
}
