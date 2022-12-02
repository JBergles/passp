package main

import (
	"bufio"
	_ "embed"
	"fmt"
	"strings"
)

//go:embed words.txt
var wordList string

func main() {
	words := loadWords()
	for _, word := range words {
		fmt.Println(word)
	}
}

func loadWords() []string {
	words := make([]string, 0, 1000)
	scanner := bufio.NewScanner(strings.NewReader(wordList))
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "#") {
			continue
		}
		words = append(words, scanner.Text())
	}
	return words
}
