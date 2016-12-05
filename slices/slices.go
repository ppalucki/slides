package main

import (
	"fmt"
)

// START OMIT
func modifySlice(s []int) {
	s[0] = -1
}

func main() {
	a := []int{42}
	fmt.Println(a)
	modifySlice(a) // pass by copy // HL
	fmt.Println(a)
}

// STOP OMIT
