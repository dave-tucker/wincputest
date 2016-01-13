package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	var mod = syscall.NewLazyDLL("kernel32.dll")
	var proc = mod.NewProc("IsProcessorFeaturePresent")
	var PF_VIRT_FIRMWARE_ENABLED uint32 = 21

	r1, r2, err := proc.Call(
		uintptr(PF_VIRT_FIRMWARE_ENABLED),
	)

	if r1 != 0 {
		fmt.Printf("Error %s, %v+", err, syscall.Errno(r1))
	}
	
	resultPtr := (*bool)(unsafe.Pointer(r2))
	
	var hasVtx *bool
	if resultPtr != nil {
		hasVtx = (*bool)(unsafe.Pointer(r2))
	}
	fmt.Printf("r1: %+v\nr2: %+v\nerr: %+v\n", r1, r2, err)
	fmt.Printf("Is VTX Enabled: %+v", hasVtx)
}
