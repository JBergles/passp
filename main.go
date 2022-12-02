package main

import (
	"bufio"
	"crypto/rand"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strings"
)

//go:embed words.txt
var wordList string

func main() {
	numWords := flag.Int("n", 12, "Number of words to generate")
	flag.Parse()
	words := loadWords()
	for i := 1; i <= *numWords; i++ {
		word, err := randWord(words)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%-20s", word)
		if i%3 == 0 {
			fmt.Print("\n")
		}
	}
	if *numWords%3 != 0 {
		fmt.Print("\n")
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

func randWord(words []string) (string, error) {
	idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(words))))
	if err != nil {
		return "", err
	}
	return words[idx.Uint64()], nil
}
