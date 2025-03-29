package omp

// #include "include/streamer.h"
import "C"
import "unsafe"

var defaultStreamer *Streamer

type Streamer struct {
	handle unsafe.Pointer
}

// NewStreamer creates a new streamer instance.
func NewStreamer() *Streamer {
	return &Streamer{}
}

// GetStreamer returns the default streamer instance.
// If no streamer is set, it will create a new one.
func GetStreamer() *Streamer {
	if defaultStreamer == nil {
		defaultStreamer = NewStreamer()
	}

	return defaultStreamer
}

// SetStreamer sets the default streamer instance.
// This is useful if you want to use a custom streamer instance.
func SetStreamer(streamer *Streamer) {
	defaultStreamer = streamer
}

// CreateDynamicObject creates a dynamic object in the world.
func (s *Streamer) CreateDynamicObject(modelid int, x float32, y float32, z float32, rx float32, ry float32, rz float32, worldid int, interiorid int, playerid int) int {
	return int(C.s_createDynamicObject(C.int(modelid), C.float(x), C.float(y), C.float(z), C.float(rx), C.float(ry), C.float(rz), C.int(worldid), C.int(interiorid), C.int(playerid)))
}

// DestroyDynamicObject destroys a dynamic object.
func (s *Streamer) DestroyDynamicObject(objectid int) {
	C.s_destroyDynamicObject(C.int(objectid))
}
