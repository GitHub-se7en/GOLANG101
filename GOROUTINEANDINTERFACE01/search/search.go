package search

import (
	"log"
	"sync"
)

//this map feed.Type and Matcher ,remember the Matcher is a interface
var getMacherByFeecType = make(map[string]Matcher)

func Run(searchTerm string) {
	//retrieve the list of feeds to search through
	//every feed has a go routine but the type of feed are different
	feeds, e := RetrieveFeeds()
	//check the error
	if e != nil {
		log.Fatal("retrieve function has encounter a problem", e)
	}
	//create a unbuffered channel to receive result
	results := make(chan *Result)

	//has to set a wait group for goroutine
	var waitGroup sync.WaitGroup

	//one group to one goroutine
	waitGroup.Add(len(feeds))

	for i, feed := range feeds {
		//log.Println(i)
		//log.Println(feed)
		matcher, exists := getMacherByFeecType[feed.Type]
		if !exists {
			matcher = getMacherByFeecType["default"]
		}

		//every
		go func(matcher2 Matcher, feed2 *Feed, i int) {
			Match(matcher2, feed2, searchTerm, results, i)
			waitGroup.Done()
		}(matcher, feed, i)
	}

	go func() {
		waitGroup.Wait()

		//the channel 一边接受一边展示还是全都弄好了再展示
		close(results)
	}()

	Display(results)
}
func Register(feedType string, matcher Matcher) {
	if _, i := getMacherByFeecType[feedType]; i {
		log.Fatalln(feedType, "Mather already registered")
	}

	log.Println("Register", feedType, "matcher")
	getMacherByFeecType[feedType] = matcher
}
