package main

import (
	"testing"
	"unicode"
)

func TestChar(t *testing.T) {
	chi := "小丑.720p.1080p.BD中英双字"
	// for _, c := range chi {
	// 	fmt.Printf("%s", rune(c))
	// }

	if IsChinese(chi) {
		t.Log("isChinese")
	}

	ascii := " Falling In Reverse - 2018 - Losing My Life - Single [AAC] [WEB] [iTunes Match]"
	if !IsChinese(ascii) {
		t.Log("IsEnglish")
	}
}

func IsChinese(str string) bool {
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}
