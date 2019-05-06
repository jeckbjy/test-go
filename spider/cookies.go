package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	//"golang.org/x/net/html"
	"github.com/gocolly/colly"
	//"github.com/gocolly/colly/debug"
	"log"
	//"strings"
)

type CookiesLevels struct {
	//Href    string `json:"-"`
	Href    string
	Level   string
	Letters string
	Words   []string
}

type CookiesPack struct {
	Name   string
	Href   string
	Levels []*CookiesLevels
}

type CookiesGame struct {
	Name  string
	Packs []*CookiesPack
}

var gCookiesGameList = make([]*CookiesGame, 0)
var gCookiesGameMap = make(map[string]*CookiesGame)
var gCookiesPackMap = make(map[string]*CookiesPack)
var gCookiesLevelMap = make(map[string]*CookiesLevels)
var gCookiesMtx = sync.Mutex{}

func runCookies() {
	log.Printf("start cookies\n")
	// 限制网页个数,用于测试
	//limitPage := 5
	limitPage := 0
	url := "https://wordcookies.info/"

	c := colly.NewCollector(
		colly.DetectCharset(),
		colly.AllowedDomains("wordcookies.info"),
		colly.Async(true))

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})

	c.OnHTML(".widget-games", func(e *colly.HTMLElement) {
		var game *CookiesGame
		//fmt.Printf("aa\n")
		e.ForEach(".c_title", func(i int, element *colly.HTMLElement) {
			name := element.ChildText("h4")
			gCookiesMtx.Lock()
			game = gCookiesGameMap[name]
			if game == nil {
				game = &CookiesGame{Name: name}
				gCookiesGameList = append(gCookiesGameList, game)
				gCookiesGameMap[name] = game
			}

			gCookiesMtx.Unlock()
			//log.Printf("%+v\n", element.ChildText("h4"))
		})

		e.ForEach(".packs", func(i int, element *colly.HTMLElement) {
			gCookiesMtx.Lock()
			defer gCookiesMtx.Unlock()

			if limitPage != 0 && len(gCookiesPackMap) >= limitPage {
				return
			}

			text := element.ChildText("a[href]")
			href := element.ChildAttr("a[href]", "href")
			//log.Printf("pack:href=%+v, text=%+v\n", href, text)
			//e.ChildAttr("a", "href")
			pack := &CookiesPack{Name: text, Href: href}
			game.Packs = append(game.Packs, pack)
			gCookiesPackMap[href] = pack

			e.Request.Visit(href)

		})
	})

	c.OnHTML(".levels", func(e *colly.HTMLElement) {
		//log.Printf("levels:%+v\n", e.Request.URL.Path)
		gCookiesMtx.Lock()
		pack := gCookiesPackMap[e.Request.URL.Path]
		gCookiesMtx.Unlock()
		if pack == nil {
			log.Printf("bad pack:%+v,%+v\n", e.Request.URL, e.Request.URL.Path)
			return
		}

		e.ForEach("li", func(i int, e1 *colly.HTMLElement) {
			if e1.Attr("class") == "ads" {
				return
			}

			level := ""
			letters := ""

			text := e1.ChildText("a")
			tokens := strings.Split(text, "Letters:")
			if len(tokens) == 2 {
				level = tokens[0]
				letters = tokens[1]
			} else if len(tokens) == 1 {
				level = tokens[0]
			} else {
				log.Printf("bad level:%+v, %+v\n", text, pack)
				return
			}

			href := e1.ChildAttr("a", "href")
			clevel := &CookiesLevels{Letters: letters, Href: href, Level: level}
			pack.Levels = append(pack.Levels, clevel)
			gCookiesMtx.Lock()
			gCookiesLevelMap[href] = clevel
			gCookiesMtx.Unlock()
			//log.Printf("level:%+v,%+v,%+v,%+v\n", level, letters, href, e.Request.URL)
			e.Request.Visit(href)
		})
	})

	c.OnHTML(".words", func(e *colly.HTMLElement) {
		log.Printf("words url:%+v", e.Request.URL.Path)
		gCookiesMtx.Lock()
		level := gCookiesLevelMap[e.Request.URL.Path]
		gCookiesMtx.Unlock()
		if level == nil {
			log.Printf("bad level:%+v\n", e.Request.URL.Path)
			return
		}

		//aa := e.DOM.Contents().Not("div")
		//tt := aa.Text()
		//log.Printf("aaaaa:%+v", tt)

		word := ""
		e.DOM.Contents().Each(func(i int, s *goquery.Selection) {
			//log.Printf("idx:%+v, text:%+v\n", i, s.Text())

			if s.Is("div") {
				// ads
				return
			}

			if s.Is("br") {
				if word != "" {
					level.Words = append(level.Words, word)
				}

				word = ""

				return
			}

			text := strings.TrimSpace(s.Text())
			if text != "" {
				word += text
			}
		})

		if word != "" {
			level.Words = append(level.Words, word)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Printf("visit:%+v\n", r.URL)
	})

	c.OnError(func(rsp *colly.Response, err error) {
		rsp.Request.Retry()
		log.Printf("error:%+v\n", err)
	})

	//c.OnResponse(func(r *colly.Response) {
	//	//log.Printf("bb:%s\n", r.Body)
	//})

	c.Visit(url)
	c.Wait()

	// output
	data, err := json.MarshalIndent(gCookiesGameList, "", "    ")
	if err == nil {
		ioutil.WriteFile("cookies.json", data, os.ModePerm)
	} else {
		log.Printf("write file fail:%+v\n", err)
	}

	log.Printf("finish cookies\n")
}
