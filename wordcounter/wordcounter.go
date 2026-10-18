package wordcounter

import (
	"strings"
)

// Изменили функцию: теперь она возвращает мапу
func WordCounter(text string) map[string]int {
	lowerText := strings.ToLower(text)
	words := strings.Fields(lowerText)
	counts := make(map[string]int, len(words)/2)

	for _, word := range words {
		// Для точности добавим очистку от знаков препинания
		cleanWord := strings.Trim(word, ".,!?;:()\"'-")
		if cleanWord != "" {
			_, ok := counts[cleanWord]

			if !ok {
				counts[cleanWord] = 1
			} else {
				counts[cleanWord]++
			}

		}
	}

	return counts
}
