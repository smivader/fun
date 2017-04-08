package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/chzyer/readline"
)

func matches(pattern, word string) bool {
	if pattern == word {
		return true
	}

	if utf8.RuneCountInString(pattern) != utf8.RuneCountInString(word) {
		return false
	}

	for {
		runePattern, size := utf8.DecodeRuneInString(pattern)
		if size == 0 {
			break
		}
		pattern = pattern[size:]

		runeWord, size := utf8.DecodeRuneInString(word)
		word = word[size:]

		if runePattern == '_' {
			continue
		}

		if runeWord != runePattern {
			return false
		}
	}

	return true
}

func main() {
	f, err := os.Open("base.dict")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	dictionary := make([]string, 0)

	buf := bufio.NewReader(f)

	for {
		l, _, err := buf.ReadLine()
		if err != nil {
			break
		}
		line := string(l)

		dictionary = append(dictionary, strings.Split(line, "/")[0])
	}

	fmt.Printf("Dictionary: %d words\n", len(dictionary))

	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		pattern, err := rl.Readline()
		if err != nil {
			break
		}
		pattern = strings.TrimSpace(pattern)

		patternRe, err := regexp.Compile(pattern)
		if err != nil {
			fmt.Printf("неправильное регулярное выражение: %s", err)
			continue
		}

		for _, word := range dictionary {
			if patternRe.MatchString(word) {
				fmt.Println(word)
			}
		}

		fmt.Println()
	}
}
