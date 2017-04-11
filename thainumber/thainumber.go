package thainumber

import (
	"strconv"
	"strings"
)

func ToThaiNumber(num int) (result string) {
	charNum := strings.Split(strconv.Itoa(num), "")
	numMap := map[string]string{
		"1": "๑",
		"2": "๒",
		"3": "๓",
		"4": "๔",
		"5": "๕",
		"6": "๖",
		"7": "๗",
		"8": "๘",
		"9": "๙",
		"0": "๐",
	}

	for _, v := range charNum {
		result += numMap[v]
	}

	return result
}
