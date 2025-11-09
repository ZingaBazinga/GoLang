package moretypes

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount23(s string) map[string]int {
	words := strings.Fields(s)

	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}
	return counts
}

func ExerciseMaps() {
	wc.Test(WordCount23)
}
