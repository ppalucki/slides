package main

import (
	"fmt"
	"reflect"
	"time"
)

type Nower interface {
	Now() time.Time
}

// first impl
type nowFunc func() time.Time

func (n nowFunc) Now() time.Time {
	return n()
}

// second impl

type s struct {
	time.Time
	name string
}

func (s s) Now() time.Time {
	return s.Time
}

var now nowFunc = time.Now

func printer(t Nower) {
	v := reflect.ValueOf(t)
	fmt.Println(v.Kind())
	fmt.Println(v.Interface())
	// fmt.Println(t.Now())
}

func main() {

	var t1 Nower = now
	var t2 Nower = s{time.Now(), "adfa"}

	printer(t1)
	printer(t2)
}
