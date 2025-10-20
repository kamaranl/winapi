//go:build windows

package winapi_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kamaranl/go_private/tools/test"
	"github.com/kamaranl/winapi"
)

var enabled = map[string]bool{
	"NewMouseInput":    true,
	"NewKeybdInput":    true,
	"NewHardwareInput": true,
	"SendInputMi":      true,
	"SendInputKi":      true,
}

func TestNewMouseInput(t *testing.T) {
	tName := "NewMouseInput"
	if !enabled[tName] {
		t.Skip(tName + test.TestsDisabled)
	}

	input := winapi.MOUSEINPUT{
		X:     400,
		Y:     600,
		Flags: winapi.MOUSEEVENTF_ABSOLUTE,
	}

	scenes := []test.Scene{
		{
			Input: input,
			Output: winapi.INPUT_Mi{
				Type: winapi.INPUT_MOUSE,
				Mi:   input,
			},
		},
	}

	for i, s := range scenes {
		t.Run(fmt.Sprintf(tName+" #%d", i), func(t *testing.T) {
			got := winapi.NewMouseInput(s.Input.(winapi.MOUSEINPUT))
			want := s.Output.(winapi.INPUT_Mi)

			if got.Type != want.Type {
				t.Errorf(test.ErrWantFGotF, want.Type, got.Type)
			}
			if !reflect.DeepEqual(got.Mi, want.Mi) {
				t.Errorf(test.ErrWantFGotF, want.Mi, got.Mi)
			}
		})
	}
}

func TestNewKeybdInput(t *testing.T) {
	tName := "NewKeybdInput"
	if !enabled[tName] {
		t.Skip(tName + test.TestsDisabled)
	}

	input := winapi.KEYBDINPUT{
		Vk:    0,
		Scan:  35,
		Flags: winapi.KEYEVENTF_SCANCODE,
	}

	scenes := []test.Scene{
		{
			Input: input,
			Output: winapi.INPUT_Ki{
				Type: winapi.INPUT_KEYBOARD,
				Ki:   input,
			},
		},
	}

	for i, s := range scenes {
		t.Run(fmt.Sprintf(tName+" #%d", i), func(t *testing.T) {
			got := winapi.NewKeybdInput(s.Input.(winapi.KEYBDINPUT))
			want := s.Output.(winapi.INPUT_Ki)

			if got.Type != want.Type {
				t.Errorf(test.ErrWantFGotF, want.Type, got.Type)
			}
			if !reflect.DeepEqual(got.Ki, want.Ki) {
				t.Errorf(test.ErrWantFGotF, want.Ki, got.Ki)
			}
		})
	}
}

func TestNewHardwareInput(t *testing.T) {
	tName := "NewHardwareInput"
	if !enabled[tName] {
		t.Skip(tName + test.TestsDisabled)
	}

	input := winapi.HARDWAREINPUT{
		Msg: uint32(winapi.WM_QUIT),
	}

	scenes := []test.Scene{
		{
			Input: input,
			Output: winapi.INPUT_Hi{
				Type: winapi.INPUT_HARDWARE,
				Hi:   input,
			},
		},
	}

	for i, s := range scenes {
		t.Run(fmt.Sprintf(tName+" #%d", i), func(t *testing.T) {
			got := winapi.NewHardwareInput(s.Input.(winapi.HARDWAREINPUT))
			want := s.Output.(winapi.INPUT_Hi)

			if got.Type != want.Type {
				t.Errorf(test.ErrWantFGotF, want.Type, got.Type)
			}
			if !reflect.DeepEqual(got.Hi, want.Hi) {
				t.Errorf(test.ErrWantFGotF, want.Hi, got.Hi)
			}
		})
	}
}
