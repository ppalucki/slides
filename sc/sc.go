package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func basic() {
	// START_BASIC OMIT
	question := "question of life"
	if question != "question of life" { // jne
		goto FAILURE
	}
	fmt.Println("42")
	goto DONE
FAILURE:
	fmt.Println("failure")
DONE:
	// STOP_BASIC OMIT
}

// pascal/c
// structured programming
func structured() {

	// START_STRUCTURED OMIT
	supercomputer := func(question string) string {
		if question != "question of life" {
			return "failure"
		}
		return "42"
	}

	answer := supercomputer("question of life")
	fmt.Println(answer)
	// STOP_STRUCTURED OMIT
}

// structured parallel programinng
// exameples: mpi/opencl, map/reduce
// expclict parallelism
// multicore Moore law
func parallel() {

	// START_PARALLEL OMIT
	supercomputer := func(question string) string {
		switch question {
		case "question o":
			return "4"
		case "f life":
			return "2"
		default:
			return "failure"
		}
	}
	answers := []string{}
	// map and some magic incantation to make it parallel
	for _, questionPart := range []string{"question o", "f life"} {
		answers = append(answers, supercomputer(questionPart))
	}
	fmt.Println(strings.Join(answers, "")) // reduce
	// STOP_PARALLEL OMIT
}

// START_v0 OMIT
func supercomputer_v0(question, answer chan string) {
	switch <-question {
	case "question o":
		answer <- "4"
	case "f life":
		answer <- "2"
	default:
		answer <- "failure"
	}
}

// STOP_v0 OMIT

func concurrent() {
	// START_CONCURRENT OMIT
	question := make(chan string)
	answers := make(chan string)
	go supercomputer_v0(question, answers)
	go supercomputer_v0(question, answers)
	question <- "question o"
	question <- "f life"
	answer := <-answers + <-answers

	fmt.Println(answer)
	// STOP_CONCURRENT OMIT
}

// io
// no io, no fun
// (you can as well program in pure functionall langugue or math)
// we have an collection an input (mutable)
func collection() {

	// not io.Reader, not streamed
	// START_COLLECTION OMIT
	supercomputer := func(question []byte) byte {
		if string(question) != "question of life" {
			return 22 /* EINVAL */
		}
		return 42
	}
	question, _ := ioutil.ReadAll(os.Stdin)
	answer := supercomputer(question)
	fmt.Println(answer)
	// STOP_COLLECTION OMIT
}

// START_FINAL_FUNC_v1 OMIT
func supercomputer_v1(question io.Reader, answer io.Writer) {
	realQuestion, _ := ioutil.ReadAll(question)
	if string(realQuestion) != "question of life" {
		io.WriteString(answer, "failure")
		return
	}
	io.WriteString(answer, "42")
}

// STOP_FINAL_FUNC_v1 OMIT

// io
// no io, no fun
func interfaces() {
	// START_INTERFACES OMIT
	supercomputer_v1(os.Stdin, os.Stderr)
	// STOP_INTERFACES OMIT
}

// and real io concurrency and abstractions
func final_works() {
	// START_FINAL_WORKS OMIT
	question := os.Stdin
	answer := os.Stderr
	go supercomputer_v1(question, answer)
	go supercomputer_v1(question, answer)
	// STOP_FINAL_WORKS OMIT
	time.Sleep(1 * time.Second)
}

// START_FINAL_FUNC_v2 OMIT
func supercomputer_v2(question io.ReadCloser, answer io.Writer) {
	realQuestion, _ := ioutil.ReadAll(question)
	defer question.Close() // HL
	if string(realQuestion) != "question of life" {
		io.WriteString(answer, "failure")
		return
	}
	io.WriteString(answer, "42")
}

// STOP_FINAL_FUNC_v2 OMIT

// but in production
func final_fails() {
	// START_FINAL_FAILS OMIT
	question, _ := os.Open("question.txt")
	go supercomputer_v2(question, os.Stderr)
	go supercomputer_v2(question, os.Stderr)
	// STOP_FINAL_FAILS OMIT
	time.Sleep(1 * time.Second)
}

func final_server() {
	// START_FINAL_SERVER  OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		supercomputer_v2(r.Body, w)
	})
	http.ListenAndServe(":8080", nil)
	// STOP_FINAL_SERVER  OMIT
}

func main() {
	// basic()
	// structured()
	// parallel()
	// concurrent()
	// collection()
	// interfaces()
	final_works()
	// final_fails()
	// for i := 0; i < 10; i++ {
	// 	final_fails()
	// 	fmt.Println("")
	// }
	// final_server()
}
