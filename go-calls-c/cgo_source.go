package main

// this include path is wrong as it should be relative to the current folder, this does not build with the go toolchain

// #include "go-calls-c/call_from_go.h"
import "C"

func CallC() {
	C.helloFromC()
}
