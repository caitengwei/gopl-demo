package main

import (
	"fmt"
	"os"
)

func main() {
	s, seq := "", ""
	// assign index to blank identifier `_` if it's not needed
	for id, arg := range os.Args[1:] {
		s += seq + arg
		seq = " "
		fmt.Printf("Index: %d\n", id)
	}
	fmt.Println(s)
}