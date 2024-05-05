package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Sample text (you can replace this with actual file reading logic)
	text := `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.`

	// Minimum character threshold for a valid word
	minCharThreshold := 3

	// Preprocess the text
	processedText := preprocessText(text)

	// Tokenize the processed text
	words := strings.Fields(processedText)

	// Filter words based on the minimum character threshold
	filteredWords := filterWords(words, minCharThreshold)

	// Count word frequencies
	wordFreq := make(map[string]int)
	for _, word := range filteredWords {
		wordFreq[word]++
	}

	// Convert word frequency map to list of word-frequency pairs
	wordFreqList := make([]wordFrequency, 0, len(wordFreq))
	for word, freq := range wordFreq {
		wordFreqList = append(wordFreqList, wordFrequency{word, freq})
	}

	// Sort word frequency list by frequency (descending order)
	sort.Slice(wordFreqList, func(i, j int) bool {
		return wordFreqList[i].frequency > wordFreqList[j].frequency
	})

	// Get the top K words
	K := 5
	topWords := getTopKWords(wordFreqList, K)

	// Print the top K words
	fmt.Printf("Top %d most common words:\n", K)
	for i, word := range topWords {
		fmt.Printf("%d. %s (%d occurrences)\n", i+1, word.word, word.frequency)
	}
}

// Structure to hold word and its frequency
type wordFrequency struct {
	word      string
	frequency int
}

// Preprocess the text (convert to lowercase and remove punctuations)
func preprocessText(text string) string {
	text = strings.ToLower(text)
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	return reg.ReplaceAllString(text, " ")
}

// Filter words based on the minimum character threshold
func filterWords(words []string, minCharThreshold int) []string {
	filtered := make([]string, 0, len(words))
	for _, word := range words {
		if len(word) >= minCharThreshold {
			filtered = append(filtered, word)
		}
	}
	return filtered
}

// Get the top K words based on word frequencies
func getTopKWords(wordFreqList []wordFrequency, K int) []wordFrequency {
	if K > len(wordFreqList) {
		K = len(wordFreqList)
	}
	return wordFreqList[:K]
}
