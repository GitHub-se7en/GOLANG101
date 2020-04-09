// Package words provides support for counting words.
// words这个包是用来提供计数功能的
package words

import "strings"

func CountWords(text string) (count int) {
	count = len(strings.Fields(text))
	return
}
