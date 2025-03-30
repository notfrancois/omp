package omp

// #include <stdlib.h>
// #include <string.h>
// #include "include/omp.h"
import "C"
import (
	"strings"
	"unicode/utf8"
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
	// Try to preserve the original string if it's valid UTF-8
	if utf8.ValidString(goStr) {
		// First attempt: direct conversion if it's valid UTF-8
		cStr := C.CString(goStr)
		return C.String{
			buf:    cStr,
			length: C.strlen(cStr),
		}
	}

	// Second attempt: try Windows-1251 encoding
	encoder := charmap.Windows1251.NewEncoder()
	encodedBytes, _, err := transform.Bytes(encoder, []byte(goStr))
	if err == nil {
		cStr := C.CString(string(encodedBytes))
		return C.String{
			buf:    cStr,
			length: C.strlen(cStr),
		}
	}

	// Fallback: clean the string by removing invalid characters
	var buf strings.Builder
	for _, r := range goStr {
		if r != utf8.RuneError && r <= 0xFFFF {
			buf.WriteRune(r)
		} else {
			buf.WriteRune('?')
		}
	}

	cleanStr := buf.String()
	cStr := C.CString(cleanStr)

	return C.String{
		buf:    cStr,
		length: C.strlen(cStr),
	}
}

func freeCString(cStr C.String) {
	C.free(unsafe.Pointer(cStr.buf))
}

// Para convertir strings de C a Go de forma segura
func safeGoString(cStr C.String) string {
	length := int(cStr.length)

	// Validación de longitud para evitar problemas de memoria
	if length < 0 || length > 4096 {
		return ""
	}

	goStr := C.GoStringN(cStr.buf, C.int(length))

	// Intenta decodificar como Windows-1251 si es necesario
	if !utf8.ValidString(goStr) {
		decoder := charmap.Windows1251.NewDecoder()
		if decodedBytes, _, err := transform.Bytes(decoder, []byte(goStr)); err == nil {
			return string(decodedBytes)
		}
	}

	return goStr
}
