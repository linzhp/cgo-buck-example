package main

// #include "call_from_go.h"
import "C"

func CallC() {
	C.helloFromC()
}
