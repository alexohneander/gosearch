package index

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

// needs to be saved in a file or database
type InvertedIndex map[string]map[string]int
type DocumentFrequency map[string]int

var Index InvertedIndex
var DocFreq DocumentFrequency
var Documents []string

type SavedIndex struct {
	Index     InvertedIndex
	DocFreq   DocumentFrequency
	Documents []string
}

func InitIndex(name string) {
	createIndex()

	// check if index as file exists
	// if not, create one and save it
	indexFilePath := "/tmp/" + name + ".db"
	if _, err := os.Stat(indexFilePath); errors.Is(err, os.ErrNotExist) {
		f, err := os.Create(indexFilePath)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
	} else {
		var savedIndex SavedIndex

		err = readStructFromFile(indexFilePath, &savedIndex)
		if err != nil {
			fmt.Println("Fehler beim Lesen:", err)
			return
		}

		Index = savedIndex.Index
		DocFreq = savedIndex.DocFreq
		Documents = savedIndex.Documents
	}
}

func createIndex() {
	index := make(InvertedIndex)
	docFreq := make(DocumentFrequency)
	var docs []string

	Documents = docs
	Index = index
	DocFreq = docFreq
}

func updateIndex(name string) {
	savedIndex := SavedIndex{
		Index:     Index,
		DocFreq:   DocFreq,
		Documents: Documents,
	}

	indexFilePath := "/tmp/" + name + ".db"
	err := writeStructToFile(indexFilePath, savedIndex)
	if err != nil {
		fmt.Println("Fehler beim Schreiben:", err)
		return
	}
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

	updateIndex("default")
}

func writeStructToFile(filename string, data interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(data)
	if err != nil {
		return fmt.Errorf("error when encoding the structs: %w", err)
	}

	err = os.WriteFile(filename, buf.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("error when writing to the file: %w", err)
	}
	return nil
}

func readStructFromFile(filename string, data interface{}) error {
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading the file: %w", err)
	}

	buf := bytes.NewBuffer(content)
	dec := gob.NewDecoder(buf)
	err = dec.Decode(data)
	if err != nil {
		return fmt.Errorf("error when decoding the structs: %w", err)
	}
	return nil
}
