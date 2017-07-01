package helper

import (
	"crypto/md5"
	"fmt"
)

func MD5BytesToString(bts []byte) string {
	return fmt.Sprintf("%x", md5.Sum(bts))
}
