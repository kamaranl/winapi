//go:build windows

// Package winapi implements [Win32 API] functions, types, and constants
// that have not yet been implemented in [syscall] or [sys.windows].
//
// [Win32 API]: https://learn.microsoft.com/en-us/windows/win32/apiindex/windows-api-list
// [sys.windows]: https://pkg.go.dev/golang.org/x/sys/windows
package winapi

// #region constants

// VK_UNASSIGNED is the last unassigned virtual key code.
const VK_UNASSIGNED = 0xE8

// SFVIDM_REFRESH is the id sent to refresh a menu/window.
const SFVIDM_REFRESH = 41504

// #endregion
// #region factories

// NewMouseInput returns an [INPUT_Mi] with the provided [MOUSEINPUT].
func NewMouseInput(mi MOUSEINPUT) (input INPUT_Mi) {
	input.Type = INPUT_MOUSE
	input.Mi = mi
	return input
}

// NewKeybdInput returns an [INPUT_Ki] with the provided [KEYBDINPUT].
func NewKeybdInput(ki KEYBDINPUT) (input INPUT_Ki) {
	input.Type = INPUT_KEYBOARD
	input.Ki = ki
	return input
}

// NewHardwareInput returns an [INPUT_Hi] with the provided [HARDWAREINPUT].
//
// Experimental: NewHardwareInput has not been tested or used internally.
func NewHardwareInput(hi HARDWAREINPUT) (input INPUT_Hi) {
	input.Type = INPUT_HARDWARE
	input.Hi = hi
	return input
}

// #endregion
// #region helpers

// toBOOL converts a Go bool to a C BOOL (int32).
func toBOOL(b bool) int32 {
	if b {
		return 1
	}

	return 0
}

// #endregion
