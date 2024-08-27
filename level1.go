package githubpack

import (
	"crypto/md5"
	"crypto/sha1"

	goutil "github.com/lialon/goutil/module_hash_test"
)

func Md5_1() [16]byte {
	return md5.Sum([]byte("123456"))
}
func Sha1_1() [20]byte {
	return sha1.Sum([]byte("123456"))
}

func Goutil() {
	goutil.Test1()
	goutil.Test2()
}
