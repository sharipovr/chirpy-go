package main

import "strings"

func helperBadWordsClean(text string) string {

	words := strings.Split(text, " ")
	for i, word := range words {
		if strings.ToLower(word) == "kerfuffle" ||
			strings.ToLower(word) == "sharbert" ||
			strings.ToLower(word) == "fornax" {
			words[i] = "****"
		}
	}

	return strings.Join(words, " ")
}
