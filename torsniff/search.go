package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blevesearch/bleve"
)

var gSearch = newSearch()

const (
	bleveFile = "index"
)

type Search struct {
	index bleve.Index
}

func newSearch() *Search {
	s := &Search{}
	return s
}

func (s *Search) Run() {
	var index bleve.Index
	var err error
	if fileExists(bleveFile) {
		index, err = bleve.Open(bleveFile)
	} else {
		mapping := bleve.NewIndexMapping()
		index, err = bleve.New(bleveFile, mapping)
	}

	if err != nil {
		fmt.Println(err)
		return
	}

	s.index = index

	fmt.Printf("search service start\n")
	http.HandleFunc("/", s.Query)
	http.ListenAndServe(":80", nil)
}

func (s *Search) Index(t *torrent) {
	if s.index == nil {
		return
	}
	data := struct {
		Name string
	}{
		Name: t.name,
	}
	s.index.Index(string(t.infohashHex), data)
}

func (s *Search) Query(w http.ResponseWriter, r *http.Request) {
	log.Printf("query:%+v", r.URL)
	match := r.URL.Query().Get("query")
	query := bleve.NewMatchQuery(match)
	search := bleve.NewSearchRequest(query)
	result, err := s.index.Search(search)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Println(result)
	data, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}

	return info.IsDir()
}
