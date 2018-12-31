package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Sorter []string

func main() {
	input := strings.ToLower(os.Args[1])
	length := len(input)

	if length < 2 {
		panic("need more than 1 char")
	}

	letters := make(map[rune]int)

	for _, letter := range []rune(input) {
		_, exists := letters[letter]

		if !exists {
			letters[letter] = count(input, letter)
		}
	}

	file, err := os.Open("words")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	output := []string{}

	for scanner.Scan() {
		text := strings.ToLower(scanner.Text())
		word := make(map[rune]int)
		skip := false
		leng := len(text)

		if leng < 2 {
			continue
		}

		for _, letter := range []rune(text) {
			_, has := word[letter]

			if !has {
				word[letter] = count(text, letter)
			}

			skip = length < leng

			if skip {
				break
			}
		}

		if skip {
			continue
		}

		wilds, has := letters[([]rune("*"))[0]]

		if !has {
			wilds = 0
		}

		for letter, c := range word {
			count, has := letters[letter]

			skip = !has || count < c

			if skip {
				wilds = wilds - (c - count)
				skip = wilds < 0
			}

			if skip {
				break
			}
		}

		if skip {
			continue
		}

		output = append(output, text)
	}

	sort.Sort(Sorter(output))

	for _, word := range output {
		fmt.Printf("%+v\t%+v\n", len(word), word)
	}
}

func count(str string, char rune) int {
	return strings.Count(str, string(char))
}

func (s Sorter) Len() int {
	return len(s)
}

func (s Sorter) Swap(i int, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func (s Sorter) Less(i int, j int) bool {
	return len(s[i]) < len(s[j])
}
