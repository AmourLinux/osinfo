package goInfo

import (
	"strings"
	"testing"
)

func TestGetInfo(t *testing.T) {
	osi, err := GetInfo()
	if err != nil {
		t.Log(err)
		t.Fail()
	}
	t.Log(osi)

	GOOsLower := strings.ToLower(osi.GoOSBig)
	if GOOsLower != osi.GoOS {
		t.Log(osi.GoOS)
		t.Fail()
	}
}
