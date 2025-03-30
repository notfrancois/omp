package omp

// #include <stdlib.h>
// #include <string.h>
// #include "include/omp.h"
import "C"
import (
	"strings"
	"unicode/utf8"
	"unsafe"
)

func newCUchar(goBool bool) C.uchar {
	if goBool {
		return 1
	}

	return 0
}

func newCString(goStr string) C.String {
	// Validate the UTF-8 string before conversion
	if !utf8.ValidString(goStr) {
		// Replace invalid sequences
		var buf strings.Builder
		for _, r := range goStr {
			if r == utf8.RuneError {
				buf.WriteRune('?')
			} else {
				buf.WriteRune(r)
			}
		}
		goStr = buf.String()
	}

	cStr := C.CString(goStr)

	return C.String{
		buf:    cStr,
		length: C.strlen(cStr),
	}
}

func freeCString(cStr C.String) {
	C.free(unsafe.Pointer(cStr.buf))
}
