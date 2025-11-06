package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a = "Hello, world!"
	bytes := []byte(a)
	fmt.Println(bytes)
	upA := uintptr(unsafe.Pointer(&a))
	upA += 1

	c := (*uint8)(unsafe.Pointer(upA))
	fmt.Println(*c)
}
