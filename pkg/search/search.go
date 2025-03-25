package search

import (
	"math"
	"sort"
)

// SearchResult stores the document and its relevance score.
type SearchResult struct {
	Document string
	Score    float64
}

// Search processes different types of queries using TF-IDF scoring.
func Search(terms []string, queryType string, index InvertedIndex, docFreq DocumentFrequency, numDocs int) []SearchResult {
	scores := make(map[string]float64)

	if queryType == "AND" {
		// Ensure all terms appear in the document (AND logic)
		for _, doc := range intersectDocs(terms, index) {
			scores[doc] = scoreDoc(terms, doc, index, docFreq, numDocs)
		}
	} else if queryType == "OR" {
		// Include any document that contains at least one of the terms (OR logic)
		for _, term := range terms {
			for doc := range index[term] {
				scores[doc] += scoreDoc([]string{term}, doc, index, docFreq, numDocs)
			}
		}
	} else {
		// Simple query - score documents based on TF-IDF for any terms
		for _, term := range terms {
			for doc := range index[term] {
				scores[doc] += scoreDoc([]string{term}, doc, index, docFreq, numDocs)
			}
		}
	}

	return rankResults(scores)
}

// Helper function to score a single document based on terms
func scoreDoc(terms []string, doc string, index InvertedIndex, docFreq DocumentFrequency, numDocs int) float64 {
	score := 0.0
	for _, term := range terms {
		tf := float64(index[term][doc])
		idf := math.Log(float64(numDocs) / float64(docFreq[term]))
		score += tf * idf
		//fmt.Printf("Score: %f64 %f64 %f64\n", tf, idf, score)
	}
	return score
}

// Helper function to intersect documents for AND logic
func intersectDocs(terms []string, index InvertedIndex) []string {
	if len(terms) == 0 {
		return nil
	}
	docs := make(map[string]bool)
	for doc := range index[terms[0]] {
		docs[doc] = true
	}
	for _, term := range terms[1:] {
		for doc := range docs {
			if _, exists := index[term][doc]; !exists {
				delete(docs, doc)
			}
		}
	}
	result := []string{}
	for doc := range docs {
		result = append(result, doc)
	}
	return result
}

// rankResults sorts the documents by score
func rankResults(scores map[string]float64) []SearchResult {
	results := make([]SearchResult, 0, len(scores))
	for doc, score := range scores {
		results = append(results, SearchResult{Document: doc, Score: score})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].Score > results[j].Score
	})
	return results
}
