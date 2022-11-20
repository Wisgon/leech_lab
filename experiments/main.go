package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	// random number
	// rand.Seed(time.Now().UnixNano())
	// aaa := rand.Intn(30)
	// fmt.Println(aaa)

	// test delete pointer
	// aaa := AAA{"fdsfsdf"}
	// bbb := &aaa
	// DeletePointer(bbb)
	// fmt.Println(bbb == nil)
	// fmt.Println(aaa)

	now := time.Now().UnixNano()
	data := []byte(fmt.Sprint(now))
	aa := fmt.Sprintf("%x", md5.Sum(data))
	fmt.Println(aa)

}

type AAA struct {
	Content string
}

func DeletePointer(aaa *AAA) {
	aaa = nil
}
