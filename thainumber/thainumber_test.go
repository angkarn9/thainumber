package thainumber

import "testing"

func TestNumber1(t *testing.T) {
	thaiNumber := ToThaiNumber(1)
	if thaiNumber != "๑" {
		t.Errorf("expect 1 to thai number is ๑, but it is %v", thaiNumber)
	}
}

func TestNumber2(t *testing.T) {
	thaiNumber := ToThaiNumber(2)
	if thaiNumber != "๒" {
		t.Errorf("expect 2 to thai number is ๒, but it is %v", thaiNumber)
	}
}

func TestNumber(t *testing.T) {
	thaiNumber := ToThaiNumber(1234567890987654321)
	if thaiNumber != "๑๒๓๔๕๖๗๘๙๐๙๘๗๖๕๔๓๒๑" {
		t.Errorf("expect 1234567890987654321 to thai number is ๑๒๓๔๕๖๗๘๙๐๙๘๗๖๕๔๓๒๑, but it is %v", thaiNumber)
	}
}
