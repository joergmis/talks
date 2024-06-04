package main

// #cgo CFLAGS: -g -Wall
// #include "test.c"
import "C"

func main() {
	println(C.add(10, 15))
}
