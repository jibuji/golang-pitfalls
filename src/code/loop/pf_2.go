package main

import (
	"fmt"
	"time"
)

type valType int

func (o *valType) Print() {
	fmt.Println(*o)
}

func main() {
	mvalus := []valType{1, 2, 3, 4}
	fmt.Println("----- pitfall outputs: -------")
	for _, val := range mvalus {
		go val.Print()
	}
	time.Sleep(time.Second)
	fmt.Println("----- corrected outputs: -------")
	for _, val := range mvalus {
		val := val
		go val.Print()
	}
	time.Sleep(time.Second)
}
