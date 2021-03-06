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

// START_BASIC OMIT
func main_() {
	question := "question of life"
	if question != "question of life" { // jne
		goto FAILURE
	}
	fmt.Println("42")
	goto DONE
FAILURE:
	fmt.Println("failure")
DONE:
}

// STOP_BASIC OMIT

// START_FUNC_STRUCTURED OMIT
func supercomputerStructured(question string) string {
	if question != "question of life" {
		return "failure"
	}
	return "42"
}

// STOP_FUNC_STRUCTURED OMIT

// pascal/c
// structured programming
func structured() {
	// START_STRUCTURED OMIT
	answer := supercomputerStructured("question of life")
	fmt.Println(answer)
	// STOP_STRUCTURED OMIT
}

// START_FUNC_PARALLEL OMIT
func supercomputerParallel(question string) string {
	switch question {
	case "question o":
		return "4"
	case "f life":
		return "2"
	default:
		return "failure"
	}
}

// STOP_FUNC_PARALLEL OMIT

// structured parallel programinng
// exameples: mpi/opencl, map/reduce
// expclict parallelism
// multicore Moore law
func parallel() {
	// START_PARALLEL OMIT

	answers := []string{}
	// map and some magic incantation to make it parallel
	for _, questionPart := range []string{"question o", "f life"} {
		answers = append(answers, supercomputerParallel(questionPart))
	}
	fmt.Println(strings.Join(answers, "")) // reduce
	// STOP_PARALLEL OMIT
}

// START_v0 OMIT
func supercomputerConcurrent(question, answer chan string) {
	switch <-question { // HL
	case "question o":
		answer <- "4" // HL
	case "f life":
		answer <- "2" // HL
	default:
		answer <- "failure" // HL
	}
}

// STOP_v0 OMIT

func concurrent() {
	// START_CONCURRENT OMIT
	question := make(chan string)
	answers := make(chan string)
	go supercomputerConcurrent(question, answers) // HL
	go supercomputerConcurrent(question, answers) // HL
	question <- "question o"
	question <- "f life"
	answer := <-answers + <-answers

	fmt.Println(answer)
	// STOP_CONCURRENT OMIT
}

// START_FUNC_COLLECTION OMIT
func supercomputerCollection(question []byte) byte {
	if string(question) != "question of life" {
		return 22 /* EINVAL */
	}
	return 42
}

// STOP_FUNC_COLLECTION OMIT

// io
// no io, no fun
// (you can as well program in pure functionall langugue or math)
// we have an collection an input (mutable)
func collection() {

	// not io.Reader, not streamed
	// START_COLLECTION OMIT
	question, _ := ioutil.ReadAll(os.Stdin)
	answer := supercomputerCollection(question)
	fmt.Println(answer)
	// STOP_COLLECTION OMIT
}

// START_FINAL_FUNC_v1 OMIT
func supercomputerInterfaces(question io.Reader, answer io.Writer) {
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
	supercomputerInterfaces(os.Stdin, os.Stderr)
	// STOP_INTERFACES OMIT
}

// and real io concurrency and abstractions
func finalWorks() {
	// START_FINAL_WORKS OMIT
	question := os.Stdin // HL
	answer := os.Stderr
	go supercomputerInterfaces(question, answer)
	go supercomputerInterfaces(question, answer)
	// STOP_FINAL_WORKS OMIT
	time.Sleep(1 * time.Second)
}

// START_FINAL_FUNC_v2 OMIT
func supercomputerInterfacesFixed(question io.ReadCloser, answer io.Writer) {
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
func finalFails() {
	// START_FINAL_FAILS OMIT
	question, _ := os.Open("question.txt") // HL
	answer := os.Stderr
	go supercomputerInterfacesFixed(question, answer)
	go supercomputerInterfacesFixed(question, answer)
	// STOP_FINAL_FAILS OMIT
	time.Sleep(1 * time.Second)
}

func finalServer() {
	// START_FINAL_SERVER  OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		supercomputerInterfacesFixed(r.Body, w)
	})
	http.ListenAndServe(":8080", nil)
	// STOP_FINAL_SERVER  OMIT
}

func main() {
	main_()
	structured()
	parallel()
	concurrent()
	collection()
	interfaces()
	finalWorks()
	finalFails()
	for i := 0; i < 10; i++ {
		finalFails()
		fmt.Println("")
	}
	finalServer()
}
