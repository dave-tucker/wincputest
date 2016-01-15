#include "windows.h"
#include <stdio.h>

int main()
{ 	
	int i = IsProcessorFeaturePresent(PF_VIRT_FIRMWARE_ENABLED);
	printf("VTX Enabled: %s", ( i== 1) ? "true" : "false");
}