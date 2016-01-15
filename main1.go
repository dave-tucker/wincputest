package main

import (
	"fmt"
	"syscall"
)

var (
	kernel32, _                  = syscall.LoadLibrary("kernel32.dll")
	isProcessorFeaturePresent, _ = syscall.GetProcAddress(kernel32, "IsProcessorFeaturePresent")
)

const (
	PF_VIRT_FIRMWARE_ENABLED = 21
)

func main() {
	r1, r2, _ := syscall.Syscall(uintptr(isProcessorFeaturePresent),
		uintptr(PF_VIRT_FIRMWARE_ENABLED),
		0,
		0,
		0)
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
