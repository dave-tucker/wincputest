package main

import (
	"fmt"
	"syscall"
)

const (
	PF_VIRT_FIRMWARE_ENABLED = 21
)

func main() {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("IsProcessorFeaturePresent")

	r1, r2, _ := proc.Call(uintptr(PF_VIRT_FIRMWARE_ENABLED))
	if r1 != 0 {
		fmt.Printf("Err: %s", syscall.Errno(r1))
	}
	resultPtr := unsafe.Pointer(r2)
	hasVTX := false
	if *resultPtr != 0 {
		hasVTX = true
	}

	fmt.Printf("VTX Enabled: %v", hasVTX)
}
