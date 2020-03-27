package search

type defaultMatcher struct {
}

func init() {
	var defaultMatcher defaultMatcher
	Register("default", defaultMatcher)
}
func (m defaultMatcher) Search(feed *Feed, searchTerm string, i int) ([]*Result, error) {
	return nil, nil
}
