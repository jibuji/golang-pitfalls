package main

import (
	"fmt"
	"time"
)

func main() {
	values := []int{1, 2, 3, 4}
	fmt.Println("----- pitfall outputs: -------")
	for _, val := range values {
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("----- correct-1 outputs: -------")
	for _, val := range values {
		val := val
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("----- correct-2 outputs: -------")
	for _, val := range values {
		go func(val int) {
			fmt.Println(val)
		}(val)
	}
	time.Sleep(time.Second)
	fmt.Println("----- correct-3 outputs: -------")
	for i := range values {
		val := values[i]
		go func() {
			fmt.Println(val)
		}()
	}
	time.Sleep(time.Second)
}
