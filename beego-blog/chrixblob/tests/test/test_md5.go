package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func GetMd5(pwd string) string {
	h := md5.New()
	h.Write([]byte(pwd))

	return hex.EncodeToString(h.Sum(nil))
}

func main() {

	md5_pwd := GetMd5("fai")
	fmt.Println(md5_pwd)

}
