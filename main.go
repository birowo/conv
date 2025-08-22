// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"unsafe"
)

func Reverse(bs []byte, i, j int) {
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i, j = i+1, j-1
	}
}

type Bs[T any] struct {
	Ptr      *T
	Len, Cap int
}

func BsAny[T any](v *T) (int, *[]byte) {
	sz := int(unsafe.Sizeof(v))
	return sz, (*[]byte)(unsafe.Pointer(&Bs[T]{v, sz, sz}))
}
func AnyBs[T any](bs *[]byte) *T {
	return (*Bs[T])(unsafe.Pointer(bs)).Ptr
}
func main() {
	fmt.Println("Hello, 世界")

	v1 := 0xabcd
	sz1, bs1 := BsAny(&v1)
	fmt.Printf("bs1: %x\n", *bs1)
	v1 = *AnyBs[int](bs1)
	fmt.Printf("v1: %x\n", v1)
	fmt.Printf("v1 type: %T\n", v1)
	Reverse(*bs1, 0, sz1-1)
	fmt.Printf("bs1 reverse: %x\n", *bs1)
	println()
	_, bs2 := BsAny(&[...]int16{-1, -2, -3, -4})
	fmt.Printf("bs2: %x\n", bs2)
	v2 := AnyBs[[4]int16](bs2)
	fmt.Printf("v2: %x\n", v2)
	fmt.Printf("v2 type: %T\n", v2)
}
