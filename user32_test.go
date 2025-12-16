//go:build windows

package winapi_test

import (
	"fmt"
	"testing"

	"github.com/kamaranl/gotools/test"
	"github.com/kamaranl/winapi"
)

// TestRealSendInputMi will right-click wherever the mouse is currently positioned.
func TestRealSendInputMi(t *testing.T) {
	tName := "SendInputMi"
	if !enabled[tName] {
		t.Skip(tName + test.TestsDisabled)
	}

	rightDown := winapi.NewMouseInput(winapi.MOUSEINPUT{
		X:         0,
		Y:         0,
		MouseData: 0,
		Flags:     winapi.MOUSEEVENTF_RIGHTDOWN,
		Time:      0,
		ExtraInfo: 0,
	})
	rightUp := rightDown
	rightUp.Mi.Flags |= winapi.MOUSEEVENTF_RIGHTUP

	input := []winapi.INPUT_Mi{rightDown, rightUp}
	scenes := []test.Scene{{Input: input, Output: nil}}

	test.Countdown(3)

	for i, s := range scenes {
		t.Run(fmt.Sprintf(tName+" #%d", i), func(t *testing.T) {
			if err := winapi.SendInput(s.Input.([]winapi.INPUT_Mi)); err != nil {
				t.Errorf(test.ErrWantFGotF, nil, err)
			}
		})
	}
}

// TestRealSendInputKi will send the letter 'h' wherever the cursor is currently
// positioned.
func TestRealSendInputKi(t *testing.T) {
	tName := "SendInputKi"
	if !enabled[tName] {
		t.Skip(tName + test.TestsDisabled)
	}

	down := winapi.NewKeybdInput(winapi.KEYBDINPUT{
		Vk:        0,
		Scan:      35,
		Flags:     winapi.KEYEVENTF_SCANCODE,
		Time:      0,
		ExtraInfo: 0,
	})
	up := down
	up.Ki.Flags |= winapi.KEYEVENTF_KEYUP

	input := []winapi.INPUT_Ki{down, up}
	scenes := []test.Scene{{Input: input, Output: nil}}

	test.Countdown(3)

	for i, s := range scenes {
		t.Run(fmt.Sprintf(tName+" #%d", i), func(t *testing.T) {
			if err := winapi.SendInput(s.Input.([]winapi.INPUT_Ki)); err != nil {
				t.Errorf(test.ErrWantFGotF, nil, err)
			}
		})
	}
}
