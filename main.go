package main

import (
	"fmt"

	"github.com/npt218/nacigo/src/collect"
)

func main() {
	fmt.Println("Start running ")
	collect.Hello()

	// Res, err := collect.Do(300)

	// fmt.Printf("Result: %v, Error: %v\n", Res, err == nil)

	collect.All()

}
