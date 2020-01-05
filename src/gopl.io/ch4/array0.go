package main

import "fmt"

var a = [3]int{1, 2, 3}

func main() {
	println(a[0])

	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
}
