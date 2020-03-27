package search

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type (
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}
	image struct {
		XMLName xml.Name `xml:"image"`
		URL     string   `xml:"url"`
		Title   string   `xml:"title"`
		Link    string   `xml:"link"`
	}
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          image    `xml:"image"`
		Item           []item   `xml:"item"`
	}
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel  `xml:"channel"`
	}
)
type rssMatcher struct{}

func init() {
	var rssMatcher rssMatcher
	Register("rss", rssMatcher)
}
func (m rssMatcher) Search(feed *Feed, searchTerm string, i int) ([]*Result, error) {
	var finalResult []*Result
	log.Printf("%d search feed type [%s] site [%s] uri [%s]\n", i, feed.Type, feed.Name, feed.URI)

	document, e := m.retrieve(feed)
	if e != nil {
		return nil, e
	}
	for _, channelItem := range document.Channel.Item {
		matched, e := regexp.MatchString(searchTerm, channelItem.Title)
		if e != nil {
			return nil, e
		}
		if matched {
			finalResult = append(finalResult, &Result{
				//先搜索channelItem里面的title，然后在搜索description
				Field:   "Title",
				Content: channelItem.Title,
			})
		}
		matched, e = regexp.MatchString(searchTerm, channelItem.Description)
		if e != nil {
			return nil, e
		}

		if matched {
			finalResult = append(finalResult, &Result{
				Field:   "Description",
				Content: channelItem.Description,
			})
		}
	}
	return finalResult, nil

}
func (m rssMatcher) retrieve(feed *Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("feed uri has no value")
	}

	//怎么把返回的数据封装起来呢？
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("http response error %s\n", resp.StatusCode)
	}

	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}
