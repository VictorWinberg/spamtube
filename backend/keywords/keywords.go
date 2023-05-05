package keywords

import (
	"math"
	"regexp"
	"sort"
	"strings"
)

// WordCount holds word and count pair
type WordCount struct {
	word  string
	count int
}

// Extract keywords from a string.
// Inspired by https://github.com/securisec/go-keywords/blob/master/keywords.go
func Extract(s string) ([]string, error) {
	var (
		results      []string
		words        []string
		matchWords   []string
		strippedHTML string
	)

	strippedHTML = s

	splitRe := regexp.MustCompile(`\s`)
	words = splitRe.Split(strings.TrimSpace(strippedHTML), -1)

	if len(words) == 0 {
		return catchErr(nil)
	}

	specialChars := "\\.|,|;|!|\\?|\\(|\\)|:|\"|\\^'|\\$|“|”|‘|’|”|<|>|–"

	for _, w := range words {
		w = regexp.MustCompile(specialChars).ReplaceAllString(w, "")
		if len(w) == 1 {
			w = regexp.MustCompile(`\-|_|@|&|#`).ReplaceAllString(w, "")
		}
		matchWords = append(matchWords, w)
	}

	for _, word := range matchWords {
		if matcher(strings.ToLower(word)) {
			results = append(results, strings.ToLower(word))
		}
	}

	results = extractMostPopularKeywords(results, 9)

	var keywords []string
	for _, words := range chunkBy(results, 3) {
		keywords = append(keywords, strings.Join(words, " "))
	}

	return keywords, nil
}

func chunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}

func catchErr(err error) ([]string, error) {
	return []string{}, err
}

func matcher(s string) bool {
	for _, w := range EnglishKeyWords {
		if w == s {
			return false
		}
	}
	return true
}

func extractMostPopularKeywords(results []string, numberOfKeyWords int) []string {
	// count same words in s
	m := make(map[string]int)
	for _, word := range results {
		m[word]++
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	maxIndex := math.Min(float64(numberOfKeyWords), float64(len(wordCounts)))
	var keywords []string

	for i := 0; i < int(maxIndex); i++ {
		keywords = append(keywords, wordCounts[i].word)
	}

	return keywords
}
