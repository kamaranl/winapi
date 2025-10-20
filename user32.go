//go:build windows

package winapi

import (
	"syscall"
	"unsafe"
)

var (
	user32                  = syscall.NewLazyDLL("user32.dll")
	procAttachThreadInput   = user32.NewProc("AttachThreadInput")
	procBlockInput          = user32.NewProc("BlockInput")
	procDispatchMessage     = user32.NewProc("DispatchMessage")
	procBringWindowToTop    = user32.NewProc("BringWindowToTop")
	procGetKeyState         = user32.NewProc("GetKeyState")
	procGetMessage          = user32.NewProc("GetMessageW")
	procGetParent           = user32.NewProc("GetParent")
	procGetWindowLongPtrW   = user32.NewProc("GetWindowLongPtrW")
	procMapVirtualKeyW      = user32.NewProc("MapVirtualKeyW")
	procMapVirtualKeyExW    = user32.NewProc("MapVirtualKeyExW")
	procPostMessageW        = user32.NewProc("PostMessageW")
	procPostThreadMessageW  = user32.NewProc("PostThreadMessageW")
	procSendInput           = user32.NewProc("SendInput")
	procSetFocus            = user32.NewProc("SetFocus")
	procSetForegroundWindow = user32.NewProc("SetForegroundWindow")
	procSetWinEventHook     = user32.NewProc("SetWinEventHook")
	procTranslateMessage    = user32.NewProc("TranslateMessage")
	procUnhookWinEvent      = user32.NewProc("UnhookWinEvent")
	procVkKeyScanExW        = user32.NewProc("VkKeyScanExW")
)

// AttachThreadInput attaches or detaches the input processing mechanism of one
// thread to that of another thread.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-attachthreadinput
func AttachThreadInput(idAttach, idAttachTo uint32, attach bool) error {
	if r1, _, err := procAttachThreadInput.Call(
		uintptr(idAttach),
		uintptr(idAttachTo),
		uintptr(toBOOL(attach)),
	); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// BlockInput blocks keyboard and mouse input events from reaching applications.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-blockinput
func BlockInput(block bool) error {
	if r1, _, err := procBlockInput.Call(uintptr(toBOOL(block))); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// BringWindowToTop brings the specified window to the top of the Z order
// (a.k.a. z-index).
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-bringwindowtotop
func BringWindowToTop(hwnd HWND) error {
	if r1, _, err := procBringWindowToTop.Call(uintptr(hwnd)); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// DispatchMessage dispatches a message to a window procedure.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-dispatchmessage
func DispatchMessage(msg MSG) {
	_, _, _ = procDispatchMessage.Call(uintptr(unsafe.Pointer(&msg)))
	// return value is intentionally ignored
}

// GetKeyState retrieves the status of the specified virtual key by specifying
// whether the key is up, down, or toggled on/off.
// It returns a pair of bools where the first bool specifies if the key is
// currently down and the second bool specifies if the key is currently toggled
// on.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getkeystate
func GetKeyState(virtKey byte) (down bool, toggled bool) {
	r1, _, _ := procGetKeyState.Call(uintptr(virtKey))
	state := int16(r1)
	down = state < 0
	toggled = (state & 0x0001) != 0

	return down, toggled
}

// GetMessage retrieves a message from the calling thread's message queue.
// It returns -1 with an error if the call fails, or 0 with no error if WM_QUIT
// is recieved. Otherwise, the return value is nonzero with no error.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getmessage
func GetMessage(msg MSG, hwnd HWND, msgFilterMin, msgFilterMax MsgId) (uintptr, error) {
	r1, _, err := procGetMessage.Call(
		uintptr(unsafe.Pointer(&msg)),
		uintptr(hwnd),
		uintptr(msgFilterMin),
		uintptr(msgFilterMax),
	)
	if int(r1) == -1 {
		if err != syscall.Errno(0) {
			return r1, err
		}

		return r1, syscall.EINVAL
	}

	return r1, nil
}

// GetParent retrieves a handle to the specified window's parent/owner.
// It returns 0 with an error if the call fails, or a [HWND] with no error on
// success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getparent
//
// Experimental: GetParent has not been tested or used internally.
func GetParent(hwnd HWND) (HWND, error) {
	r1, _, err := procGetParent.Call(uintptr(hwnd))
	if r1 == 0 {
		if err != syscall.Errno(0) {
			return 0, err
		}

		return 0, syscall.EINVAL
	}

	return HWND(r1), nil
}

// GetWindowLongPtrW retrieves information about the specified window.
// It returns 0 with an error if the call fails, or the requested value with no
// error on success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-getwindowlongptrw
//
// Experimental: GetWindowLongPtrW has not been tested or used internally.
func GetWindowLongPtrW(hwnd HWND, index GWL) (uintptr, error) {
	r1, _, err := procGetWindowLongPtrW.Call(
		uintptr(hwnd),
		uintptr(index),
	)
	if r1 == 0 {
		if err != syscall.Errno(0) {
			return 0, err
		}

		return 0, syscall.EINVAL
	}

	return r1, nil
}

// MapVirtualKeyW translates a virtual-key code into a scan code or character
// value, or translates a scan code into a virtual-key code.
// It returns 0 with an error if the call fails, or the translated key code with
// no error on success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeyw
//
// Experimental: MapVirtualKeyW has not been tested or used internally.
func MapVirtualKeyW(code uint32, mapType MapVKType) (uint32, error) {
	r1, _, err := procMapVirtualKeyW.Call(
		uintptr(code),
		uintptr(mapType),
	)
	if r1 == 0 {
		return 0, err
	}

	return uint32(r1), nil
}

// MapVirtualKeyExW translates a virtual-key code into a scan code or character
// value, or translates a scan code into a virtual-key code. It translates the
// codes using the input language and an input locale identifer.
// It returns 0 with an error if the call fails, or the translated key code with
// no error on success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-mapvirtualkeyexw
func MapVirtualKeyExW(code uint32, mapType MapVKType, hkl Handle) (uint32, error) {
	r1, _, err := procMapVirtualKeyExW.Call(
		uintptr(code),
		uintptr(mapType),
		uintptr(hkl),
	)
	if r1 == 0 {
		return 0, err
	}

	return uint32(r1), nil
}

// PostMessageW posts a message in the message queue for the specified window
// and returns without waiting for the window's thread to process the message.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postmessagew
func PostMessageW(hwnd HWND, msg MsgId, wParam, lParam uintptr) error {
	if r1, _, err := procPostMessageW.Call(
		uintptr(hwnd),
		uintptr(msg),
		wParam,
		lParam,
	); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// PostThreadMessageW posts a message to the message queue of the specified
// thread and returns without waiting for the thread to process the message.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-postthreadmessagew
func PostThreadMessageW(idThread uint32, msg MsgId, wParam, lParam uintptr) error {
	if r1, _, err := procPostThreadMessageW.Call(
		uintptr(idThread),
		uintptr(msg),
		wParam,
		lParam,
	); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// SendInput synthesizes keystrokes, mouse motions, and button clicks through
// the provided inputs. SendInput can only send a slice of one type at a time.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-sendinput
func SendInput[T INPUT_Mi | INPUT_Ki | INPUT_Hi](inputs []T) error {
	if len(inputs) == 0 {
		return nil
	}

	if r1, _, err := procSendInput.Call(
		uintptr(len(inputs)),
		uintptr(unsafe.Pointer(&inputs[0])),
		uintptr(unsafe.Sizeof(inputs[0])),
	); r1 == 0 {
		if err != syscall.Errno(0) {
			return err
		}

		return syscall.EINVAL
	}

	return nil
}

// SetFocus sets the keyboard focus to the specified window, as long as the
// window is attached to the calling thread's message queue.
// It returns 0 with an error if the call fails, or a [HWND] with no error on
// success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setfocus
func SetFocus(hwnd HWND) (HWND, error) {
	r1, _, err := procSetFocus.Call(uintptr(hwnd))
	if r1 == 0 {
		if err != syscall.Errno(0) {
			return 0, err
		}

		return 0, syscall.EINVAL
	}

	return HWND(r1), nil
}

// SetForegroundWindow brings the thread that created the specified window into
// the foreground and activates the window.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setforegroundwindow
func SetForegroundWindow(hwnd HWND) error {
	if r1, _, err := procSetForegroundWindow.Call(uintptr(hwnd)); r1 == 0 {
		return err
	}

	return nil
}

// SetWinEventHook sets an event hook function for a range of events.
// It returns 0 with an error if the call fails, or a [Handle] with no error on success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-setwineventhook
func SetWinEventHook(eventMin WEvent, eventMax WEvent, hmodWinEventProc, pfnWinEventProc uintptr, idProcess, idThread uint32, dwFlags WEFlags) (Handle, error) {
	r1, _, err := procSetWinEventHook.Call(
		uintptr(eventMin),
		uintptr(eventMax),
		uintptr(hmodWinEventProc),
		uintptr(pfnWinEventProc),
		uintptr(idProcess),
		uintptr(idThread),
		uintptr(dwFlags),
	)
	if r1 == 0 {
		if err != syscall.Errno(0) {
			return 0, err
		}

		return 0, syscall.EINVAL
	}

	return Handle(r1), nil
}

// TranslateMessage translates virtual-key messages into character messages.
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-translatemessage
func TranslateMessage(msg MSG) error {
	if r1, _, err := procTranslateMessage.Call(uintptr(unsafe.Pointer(&msg))); r1 == 0 {
		return err
	}

	return nil
}

// UnhookWinEvent removes an event hook function created by a previous call to
// [SetWinEventHook].
// It returns an error if the call fails.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-unhookwinevent
func UnhookWinEvent(winEventHook Handle) error {
	if r1, _, err := procUnhookWinEvent.Call(uintptr(winEventHook)); r1 == 0 {
		return err
	}

	return nil
}

// VkKeyScanExW translates a character to the corresponding virtual-key code and
// shift state. It translates the character using the input language and
// physical keyboard layout identifed by the input locale identifier.
// It returns (0,0) with an error if the call fails, or the translated key code
// and shift state with no error on success.
//
// See: https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-vkkeyscanexw
func VkKeyScanExW(ch int16, hkl Handle) (code byte, shift byte, err error) {
	r1, _, err := procVkKeyScanExW.Call(
		uintptr(ch),
		uintptr(hkl),
	)
	vkShift := int16(r1)
	if vkShift == -1 {
		return 0, 0, err
	}

	code = byte(vkShift & 0xFF)
	shift = byte((vkShift >> 8) & 0xFF)

	return code, shift, nil
}
