package main

import (
	"fmt"
	"log"
	"os"

	"github.com/blevesearch/bleve"
	"github.com/yanyiwu/gojieba"

	_ "jieba/jieba"
)

const (
	INDEX_DIR = "gojieba.bleve"
)

// 测试bleve的中文分词功能
func main() {
	os.RemoveAll(INDEX_DIR)
	mapping := bleve.NewIndexMapping()
	err := mapping.AddCustomTokenizer(
		"gojieba",
		map[string]interface{}{
			"dictpath":     gojieba.DICT_PATH,
			"hmmpath":      gojieba.HMM_PATH,
			"userdictpath": gojieba.USER_DICT_PATH,
			"type":         "gojieba",
		},
	)
	if err != nil {
		log.Panicf("AddCustomTokenizer fail,err=%+v", err)
	}

	err = mapping.AddCustomAnalyzer(
		"gojieba",
		map[string]interface{}{
			"type":      "gojieba",
			"tokenizer": "gojieba",
		},
	)
	if err != nil {
		log.Panicf("AddCustomAnalyzer fail, err=%+v", err)
	}
	mapping.DefaultAnalyzer = "gojieba"

	index, err := bleve.New(INDEX_DIR, mapping)
	if err != nil {
		fmt.Println(err)
		return
	}

	messages := []struct {
		Id   string
		Body string
	}{
		{
			Id:   "1",
			Body: "你好",
		},
		{
			Id:   "2",
			Body: "世界",
		},
		{
			Id:   "3",
			Body: "亲口",
		},
		{
			Id:   "4",
			Body: "交代",
		},
		{
			Id:   "5",
			Body: "[ねぎしおめろん] ホワイトブロンド留学生のメス堕ち♂ホームステイ [中国翻訳]",
		},
		{
			Id:   "6",
			Body: "Movavi Video Editor 14 Plus 14.0.0 RePack by KpoJIuK",
		},
		{
			Id:   "7",
			Body: "Gorjachie.mamochki.2018.WEB-DLRip.ELEKTRI4KA.avi",
		},
	}

	for _, msg := range messages {
		if err := index.Index(msg.Id, msg); err != nil {
			log.Panic(err)
		}
	}

	querys := []string{
		// "你好世界",
		// "亲口交代",
		"Movavi",
	}

	for _, q := range querys {
		req := bleve.NewSearchRequest(bleve.NewQueryStringQuery(q))
		// req.Highlight = bleve.NewHighlight()
		res, err := index.Search(req)
		if err != nil {
			panic(err)
		}
		for _, hit := range res.Hits {
			fmt.Printf("hitid,%+v\n", hit.ID)
		}
		fmt.Println(res)
	}
}
