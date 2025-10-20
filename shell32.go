//go:build windows

package winapi

import "syscall"

var (
	shell32            = syscall.NewLazyDLL("shell32.dll")
	procSHChangeNotify = shell32.NewProc("SHChangeNotify")
)

// SHChangeNotify notifies the system of an event, by eventId, that an
// application has performed.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/shlobj_core/nf-shlobj_core-shchangenotify
//
// Experimental: SHChangeNotify has not been tested or used internally.
func SHChangeNotify(eventId SHCNEvent, flags SHCNFlags, items ...uintptr) {
	switch len(items) {
	case 0:
		items = []uintptr{0, 0}
	case 1:
		items = append(items, 0)
	}

	_, _, _ = procSHChangeNotify.Call(
		uintptr(eventId),
		uintptr(flags),
		uintptr(items[0]),
		uintptr(items[1]),
	)
}
