package index

import (
	"bufio"
	"strings"
)

// needs to be saved in a file or database
type InvertedIndex map[string]map[string]int
type DocumentFrequency map[string]int

var Index InvertedIndex
var DocFreq DocumentFrequency
var Documents []string

func CreateIndex() {
	index := make(InvertedIndex)
	docFreq := make(DocumentFrequency)
	var docs []string

	Documents = docs
	Index = index
	DocFreq = docFreq
}

func AddDocToIndex(url string, content string) {
	Documents = append(Documents, url)

	reader := strings.NewReader(content)

	seenTerms := make(map[string]bool) // Track terms in this document
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := strings.ToLower(strings.Trim(scanner.Text(), ",.!?"))

		if Index[word] == nil {
			Index[word] = make(map[string]int)
		}
		Index[word][url]++

		if !seenTerms[word] {
			DocFreq[word]++
			seenTerms[word] = true
		}
	}

}
