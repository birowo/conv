package main

import (
	"fmt"
	"unsafe"

	"github.com/birowo/conv"
)

func main() {
	v3 := int32(0xcdef)
	sz3, bs3 := conv.BsAny(&v3)
	fmt.Printf("bs1: %x\n", *bs3)
	v3 = *conv.AnyBs[int32](bs3)
	fmt.Printf("v1: %x\n", v3)
	fmt.Printf("v1 type: %T\n", v3)
	conv.Reverse(*bs3, 0, sz3-1)
	fmt.Printf("bs1 reverse: %x\n", *bs3)
	println()
	v1 := 0xabcd
	sz1, bs1 := conv.BsAny(&v1)
	fmt.Printf("bs1: %x\n", *bs1)
	v1 = *conv.AnyBs[int](bs1)
	fmt.Printf("v1: %x\n", v1)
	fmt.Printf("v1 type: %T\n", v1)
	conv.Reverse(*bs1, 0, sz1-1)
	fmt.Printf("bs1 reverse: %x\n", bs1)
	println()
	const four = 4
	_, bs2 := conv.BsAny(&[four]int16{-257, -258, -259, -260})
	fmt.Printf("bs2: %x\n", bs2)
	v2 := conv.AnyBs[[four]int16](bs2)
	fmt.Printf("v2: %x\n", v2)
	fmt.Printf("v2 type: %T\n", v2)
	for i := range uintptr(four) {
		sz2 := unsafe.Sizeof(v2[0])
		j := i * sz2
		conv.Reverse(*bs2, j, j+sz2-1)
	}
	fmt.Printf("reverse: %x\n", bs2)
}
