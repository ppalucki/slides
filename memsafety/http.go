// +build OMIT

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {

	resp, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}

	// read/copy body
	n, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("n = ", n)
	err = resp.Body.Close()
	if err != nil {
		panic(err)
	}

	err = resp.Write(os.Stdout)
	if err != nil {
		panic(err)
	}
}
