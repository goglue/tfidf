package tfidf

import (
	"math"
	"strings"
)

// New receives the document string, the document storage and returns the weigher structure
func New(ds DocumentStorage) *Weigher {
	return &Weigher{ds}
}

// DocumentStorage encapsulates the functionality that is needed to perform the algorithm
type DocumentStorage interface {
	// DocumentsWith receives a t term parameter and returns an unsigned integer of the documents count containing t
	DocumentsWith(t string) uint
	// Documents returns the total amount of documents within the storage
	Documents() uint
}

// Weigher encapsulates the given document and the document storage where these values need to be given
type Weigher struct {
	// d the given document string
	ds DocumentStorage
}

// Score calculates the TF-IDF of a given document string. The string should contain words separated by one space only
func (w *Weigher) Score(d string) map[string]float64 {
	terms := strings.Split(d, " ")

	// tf terms frequencies within the given document
	tf := make(map[string]int)
	// tt total terms count within a given document
	tt := len(terms)

	for i := 0; i < tt; i++ {
		tf[terms[i]]++
	}

	// tft tf(t) term frequency of t
	tfidf := make(map[string]float64, len(tf))

	for term, freq := range tf {
		tft := float64(freq) / float64(tt)
		dwt := float64(w.ds.DocumentsWith(term))

		var idf float64
		if 0 == dwt {
			idf = 0
		} else {
			idf = math.Log10(
				float64(w.ds.Documents()) / dwt,
			)
		}
		tfidf[term] = tft * idf
	}

	return tfidf
}
