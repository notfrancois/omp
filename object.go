package main

import "unsafe"

type Object struct {
	handle       unsafe.Pointer
	playerHandle unsafe.Pointer
}

func NewObject() *Object {
	return &Object{}
}

func NewPlayerObject() *Object {
	return &Object{}
}
