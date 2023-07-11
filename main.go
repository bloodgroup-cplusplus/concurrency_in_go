package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {
	go printSomething("this is a string")

	time.Sleep(1 * time.Second) // this is not a good way to do it

	printSomething("This is the second thing to be printed")
}

