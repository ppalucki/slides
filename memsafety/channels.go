package main

import (
	"fmt"
)

func printer(c chan int) {
	fmt.Println(<-c)
}

func gen(c chan int) {
	c <- 1
	close(c)
}

func main() {

	ch := make(chan int)
	go gen(ch)
	go gen(ch)
	printer(ch)
}
