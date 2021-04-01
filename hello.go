package main

import (
	"fmt"

	"rsc.io/quote"

	"time"

	"example.com/hello/our_packages/add"
	"example.com/hello/our_packages/integer"
)

func main() {
	fmt.Println("hello, World!")
	fmt.Println(quote.Go())
	fmt.Println("The Time is ", time.Now());
	fmt.Println(add.Add(8,9))
	integer.Output()
}