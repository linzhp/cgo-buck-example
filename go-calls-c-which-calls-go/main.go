package main

// #include "call_go.h"
import "C"

func main() {
	C.printSomethingFromGo(1, 2)
}