package omp

// #include <stdlib.h>
// #include <string.h>
// #include "include/omp.h"
import "C"
import (
	"fmt"
	"unsafe"

	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

func newCUchar(goBool bool) C.uchar {
	if goBool {
		return 1
	}

	return 0
}

func newCString(goStr string) C.String {
	encoder := charmap.ISO8859_1.NewEncoder()
	encodedBytes, _, err := transform.Bytes(encoder, []byte(goStr))
	if err == nil {
		cStr := C.CString(string(encodedBytes))
		return C.String{
			buf:    cStr,
			length: C.strlen(cStr),
		}
	}

	return C.String{
		buf:    C.CString(goStr),
		length: C.strlen(C.CString(goStr)),
	}
}

func freeCString(cStr C.String) {
	C.free(unsafe.Pointer(cStr.buf))
}

func logStringBytes(name string, s string) {
	fmt.Print(name + ": ")
	for _, b := range []byte(s) {
		fmt.Printf("%02x ", b)
	}
	fmt.Println()
}
