package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// random number
	rand.Seed(time.Now().UnixNano())
	aaa := rand.Intn(30)
	fmt.Println(aaa)
}
