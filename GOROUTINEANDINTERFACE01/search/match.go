//我需要定义一个接口，然后另外两个类型实现它
package search

import "log"

type Result struct {
	Field   string
	Content string
}

//使用大写开头，
// 按照之前的理解的话，接口也是一种类型，实现这个接口，就是转成这个类型
type Matcher interface {
	Search(feed *Feed, searchTerm string, i int) ([]*Result, error)
}

func Match(matcher Matcher, feed *Feed, searchTerm string, chanResult chan<- *Result, i int) {
	searchResults, e := matcher.Search(feed, searchTerm, i)
	if e != nil {
		return
	}
	for _, result := range searchResults {
		chanResult <- result
	}
}
func Display(results chan *Result) {
	//unbuffered channel has no ability to store data
	//so 一边产生，一边展示
	for result := range results {
		log.Printf("%s:\n%s\n\n", result.Field, result.Content)
	}
}
