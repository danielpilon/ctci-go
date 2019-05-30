//+build !test

package main

import (
	"fmt"
	"os"
)

func main() {
	s := os.Args[1]
	fmt.Printf("with map: %v\n", isUniqueWithMap(s))
	fmt.Printf("with slice: %v\n", isUniqueWithSlice(s))
	fmt.Printf("with bit vector: %v\n", isUniqueWithBitVector(s))
}
