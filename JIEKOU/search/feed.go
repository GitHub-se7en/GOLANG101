package search

import (
	"encoding/json"
	"os"
)

const datafile = "JIEKOU/search/data.json"

type Feed struct {
	Name string `json:"site"`
	URI  string `json:"link"`
	Type string `json:"type"`
}

func RetrieveFeeds() ([]*Feed, error) {
	file, e := os.Open(datafile)
	if e != nil {
		return nil, e
	}
	//defer关键字的作用是在函数返回的时候，保证这个函数一定会执行
	defer file.Close()

	var feeds []*Feed
	e = json.NewDecoder(file).Decode(&feeds)
	return feeds, e
}
