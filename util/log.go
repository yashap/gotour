package util

import (
	"fmt"
)

// Log prints the value and type of each value given to it
func Log(vs ...interface{}) {
	lastIndex := len(vs) - 1
	for i, v := range vs {
		fmt.Printf("%v: %T", v, v)
		if i == lastIndex {
			fmt.Println()
		} else {
			fmt.Print(" ; ")
		}
	}
}

// LogIntSlice prints the length, capacity and value of a slice of ints
func LogIntSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// LogByteSlice prints the length, capacity and value of a slice of bytes
func LogByteSlice(s []byte) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
