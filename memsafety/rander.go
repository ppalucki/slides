package main

import (
	"fmt"
	"math/rand"
	"time"
)

var defaultRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func test(index int) {
	for i := 0; i < 1000*10; i++ {
		_ = defaultRand.Int63()
	}
	fmt.Println("done", index)
}

func main() {
	for i := 0; i < 10; i++ {
		go test(i)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("exit")
}
