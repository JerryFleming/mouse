package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

func main() {
	var mod = syscall.NewLazyDLL("user32.dll")
	var mouse = mod.NewProc("SwapMouseButton")
	var proc = mod.NewProc("MessageBoxTimeoutW")
	var MB_TOPMOST = 0x40000
	var MB_ICON_INFO = 0x40
	var MB_OK = 0x0
	var mode string
	ret, _, _ := mouse.Call(1)
	if ret != 0 {
		ret, _, _ = mouse.Call(0)
		mode = "right"
	} else {
		mode = "left"
	}

	ret, _, _ = proc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(fmt.Sprintf("Switched to %s handed!", mode)))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("Mouse"))),
		uintptr(MB_TOPMOST|MB_OK|MB_ICON_INFO),
		0, 3000)
	fmt.Printf("Return: %s\n", mode)
}
