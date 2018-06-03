// Switch mouse buttons on Windows.
// This program is for windows only to switch mouse buttons at a click.
package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	// Always on top.
	MB_TOPMOST = 0x40000
	// Icon style.
	MB_ICON_INFO = 0x40
	// Button to show.
	MB_OK = 0x0
)

// User module.
var mod = syscall.NewLazyDLL("user32.dll")

// Show a dialog with timeout.
func alert(title string, content string) uintptr {
	var proc = mod.NewProc("MessageBoxTimeoutW")
	ret, _, _ := proc.Call(0,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))),
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(content))),
		uintptr(MB_TOPMOST|MB_OK|MB_ICON_INFO),
		0, 3000)
	return ret
}

// The entry point. Calling this will switch mouse button.
func main() {
	var mouse = mod.NewProc("SwapMouseButton")
	var mode string
	ret, _, _ := mouse.Call(1)
	if ret != 0 {
		ret, _, _ = mouse.Call(0)
		mode = "right"
	} else {
		mode = "left"
	}

	ret = alert("mouse", fmt.Sprintf("Switched to %s handed!", mode))
	fmt.Printf("Return: %s\n", mode)
}
