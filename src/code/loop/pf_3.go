package main

import "fmt"

func main() {
	var out []*int
	fmt.Println("----- pitfall outputs: -------")
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])

	out = nil
	fmt.Println("----- corrected outputs: -------")
	for i := 0; i < 3; i++ {
		i := i
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}
