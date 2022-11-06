package main

import (
	"fmt"
)

func main() {
	var out [][]int
	fmt.Println("----- pitfall outputs: -------")
	for _, i := range [][1]int{{1}, {2}, {3}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
	out = nil
	fmt.Println("----- corrected-1 outputs: -------")
	for _, i := range [][1]int{{1}, {2}, {3}} {
		i := i // array copy
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)

	out = nil
	fmt.Println("----- corrected-2 outputs: -------")
	//use slice instead of array as the iterator value
	for _, i := range [][]int{{1}, {2}, {3}} {

		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
}
