package omp

// #include <stdlib.h>
// #include <string.h>
// #include "include/omp.h"
import "C"
import (
	"unicode/utf8"
	"unsafe"
)

func newCUchar(goBool bool) C.uchar {
	if goBool {
		return 1
	}

	return 0
}

func newCString(s string) *C.char {
	// Validate the string is valid UTF-8
	bs := []byte(s)
	for len(bs) > 0 {
		r, size := utf8.DecodeRune(bs)
		if r == utf8.RuneError {
			// Replace invalid characters with '?'
			r = '?'
		}
		bs = bs[size:]
	}

	// Convert the string to a C string while keeping UTF-8 characters
	cstr := C.CString(s)
	return cstr
}

func freeCString(cStr C.String) {
	C.free(unsafe.Pointer(cStr.buf))
}
