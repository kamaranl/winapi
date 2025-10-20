//go:build windows

package winapi

import (
	"syscall"
)

var (
	kernel32             = syscall.NewLazyDLL("kernel32.dll")
	procAttachConsole    = kernel32.NewProc("AttachConsole")
	procAllocConsole     = kernel32.NewProc("AllocConsole")
	procFreeConsole      = kernel32.NewProc("FreeConsole")
	procGetConsoleWindow = kernel32.NewProc("GetConsoleWindow")
	procSetStdHandle     = kernel32.NewProc("SetStdHandle")
)

// AllocConsole creates a new console for the calling process.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/console/allocconsole
func AllocConsole() error {
	if r1, _, err := procAllocConsole.Call(); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// AttachConsole attaches the calling process to the console of another process
// specified by pid.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/console/attachconsole
func AttachConsole(pid ACPId) error {
	if r1, _, err := procAttachConsole.Call(uintptr(pid)); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// FreeConsole detaches the calling process from its console.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/console/freeconsole
func FreeConsole() error {
	if r1, _, err := procFreeConsole.Call(); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// GetConsoleWindow retrieves the window handle of the console associated with
// the calling process.
// It returns a [Handle] representing the console window.
//
// See: https://learn.microsoft.com/en-us/windows/console/getconsolewindow
//
// Experimental: GetConsoleWindow has not been tested or used internally.
func GetConsoleWindow() Handle {
	r1, _, _ := procGetConsoleWindow.Call()
	return Handle(r1)
}

// SetStdHandle sets the handle for a standard device (input, output, or error).
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/console/setstdhandle
func SetStdHandle(stdHndl HSTDIO, fd uintptr) error {
	if r1, _, err := procSetStdHandle.Call(
		uintptr(stdHndl),
		fd,
	); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}
