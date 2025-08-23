package conv

import "unsafe"

func Reverse(bs []byte, i, j uintptr) {
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i, j = i+1, j-1
	}
}

type bsAny[T any] struct {
	Ptr      *T
	Len, Cap uintptr
}

func BsAny[T any](v *T) (uintptr, *[]byte) {
	sz := unsafe.Sizeof(*v)
	return sz, (*[]byte)(unsafe.Pointer(&bsAny[T]{v, sz, sz}))
}
func AnyBs[T any](bs *[]byte) *T {
	return (*bsAny[T])(unsafe.Pointer(bs)).Ptr
}
