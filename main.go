package main

import (
	"fmt"
	"time"
)

func printSomething(s string) {
	fmt.Println(s)
}

func main() {

	go printSomething("This is the first thing to be printed!")
	// execute in its own go routine
	//
	//this executes very fast
	// this is not a good way to do it
	// in the next lession we will see something called as weight groups

	time.Sleep(1 * time.Second)
	printSomething("This is the second thing to be printed")
}
