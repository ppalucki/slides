// +build OMIT

package main

import "time"

func whenFunc(nowFunc func() time.Time) {
	_ = nowFunc()
}

func when(t time.Time) {
	_ = t
}

func main() {

	N := 10000
	now := time.Now()
	slice := []time.Time{now}
	mape := map[int]time.Time{0: now}

	myNow0 := time.Now()
	for i := 0; i < N; i++ {
		go when(myNow0)
	}

	myNow1 := func() time.Time {
		return now
	}
	for i := 0; i < N; i++ {
		go whenFunc(myNow1)
	}

	myNow2 := func() time.Time {
		return slice[0]
	}
	for i := 0; i < N; i++ {
		go whenFunc(myNow2)
	}

	myNow3 := func() time.Time {
		return mape[0]
	}
	for i := 0; i < N; i++ {
		go whenFunc(myNow3)
	}

	time.Sleep(10 * time.Second)

}
