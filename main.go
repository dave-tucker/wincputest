package main

/*
#include "windows.h"

int hasVtx()
{
	return IsProcessorFeaturePresent(PF_VIRT_FIRMWARE_ENABLED);
}
*/
import "C"

import "fmt"

func main() {
	hasVTX := C.hasVtx() != 0
	fmt.Printf("VTX Enabled: %v", hasVTX)
}
