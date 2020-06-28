package main

import (
	"testing"

	"github.com/abadojack/whatlanggo"
)

func TestLang(t *testing.T) {
	str1 := "小丑.720p.1080p.BD中英双字"
	str2 := "Falling In Reverse - 2018 - Losing My Life - Single [AAC] [WEB] [iTunes Match]"
	str3 := "[ねぎしおめろん] ホワイトブロンド留学生のメス堕ち♂ホームステイ [中国翻訳].zip"
	info1 := whatlanggo.Detect(str1)
	info2 := whatlanggo.Detect(str2)
	info3 := whatlanggo.Detect(str3)
	t.Log(info1.Lang, whatlanggo.Scripts[info1.Script])
	t.Log(info2.Lang, whatlanggo.Scripts[info2.Script])
	t.Log(info3.Lang, whatlanggo.Scripts[info3.Script])

}
