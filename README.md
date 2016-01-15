Reliable VTX Detection for Windows
==================================

`coreinfo -v` only reports CPU capabilties, not if the BIOS is enabled
The `IsProcesserFeatureEnabled` SysInfo API gives us that information if given the right encouragement!
Namely, `PF_VIRT_FIRMWARE_ENABLED`...

## Doing it in C

Requires Mingw64
```
$ gcc main.c -o main.exe
$ main.exe
VTX Enabled: false
```
## Doing it in Go

This is much harder than it needs to be.
According to the Golang wiki there are numerous ways to do this.
I've created `main1.go`, `main2.go`, `main3.go` as examples.
None of them work

`main.go` is the only way to make this work so far, and it requires Cgo.
