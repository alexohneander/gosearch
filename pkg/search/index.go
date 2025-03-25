package search

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// needs to be saved in a file or database
type InvertedIndex map[string]map[string]int
type DocumentFrequency map[string]int

var Index InvertedIndex
var DocFreq DocumentFrequency
var Files []string

// BuildIndex reads files and builds an inverted index.
func BuildIndex(files []string) (InvertedIndex, DocumentFrequency, error) {
	index := make(InvertedIndex)
	docFreq := make(DocumentFrequency)

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, nil, err
		}
		defer f.Close()

		seenTerms := make(map[string]bool) // Track terms in this document
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			word := strings.ToLower(strings.Trim(scanner.Text(), ",.!?"))

			if index[word] == nil {
				index[word] = make(map[string]int)
			}
			index[word][file]++

			if !seenTerms[word] {
				docFreq[word]++
				seenTerms[word] = true
			}
		}
	}
	return index, docFreq, nil
}

func TestIndex() {
	// Index files
	files := []string{"data/doc1.txt", "data/doc2.txt", "data/doc3.txt"}

	index, docFreq, err := BuildIndex(files)
	if err != nil {
		log.Fatalf("Error building index: %v", err)
	}

	Files = files
	Index = index
	DocFreq = docFreq
}
