package main

import "syscall"
import "fmt"

const (
	PF_VIRT_FIRMWARE_ENABLED = 21
)

func main() {
	dll := syscall.MustLoadDLL("kernel32.dll")
	f := dll.MustFindProc("IsProcessorFeaturePresent")
	r1, r2, _ := f.Call(uintptr(PF_VIRT_FIRMWARE_ENABLED))
	if r1 != 0 {
		// handle this properly
		fmt.Printf("Err: %s", syscall.Errno(r1))
	}
	resultPtr := unsafe.Pointer(r2)
	hasVTX := false
	if *resultPtr != 0 {
		hasVTX = true
	}

	fmt.Printf("VTX Enabled: %v", hasVTX)
}
