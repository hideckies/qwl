package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
)

const (
	VERSION = "0.1.0"
	USAGE   = `QWL - Clipboard to Wordlist

qwl generates a wordlist from your clipboard text.

USAGE:
  1. Copy text from somewhere.
  2. Run "qwl" in terminal. A new wordlist will be generated.

OPTIONS:
  -min, --min    Specify the minimum word length
  -max, --max    Specify the maximum word length
  -h, --help     Print the usage of qwl
  -v, --version  Print the version of qwl
`
)

func main() {
	var help bool
	var version bool
	var min int
	var max int

	flag.IntVar(&min, "min", 3, "Specify the minimum word length")
	flag.IntVar(&max, "max", 50, "Specify the maximum word length")
	flag.BoolVar(&help, "help", false, "Print the usage of qwl")
	flag.BoolVar(&help, "h", false, "Print the usage of qwl")
	flag.BoolVar(&version, "v", false, "Print the version of qwl")
	flag.BoolVar(&version, "version", false, "Print the version of qwl")
	flag.Parse()

	if help {
		fmt.Print(USAGE)
		return
	}
	if version {
		fmt.Printf("qwl v%s\n", VERSION)
		return
	}

	s, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	re := regexp.MustCompile(`\s`)

	wordMap := make(map[string]bool)
	words := strings.FieldsFunc(s, split)
	// Store words to wordMap
	if len(words) > 0 {
		for _, word := range words {
			if word != "" && !re.MatchString(word) && len(word) >= min && len(word) <= max {
				if wordMap[word] {
					continue
				}
				wordMap[word] = true
			}
		}
	}

	for k, _ := range wordMap {
		fmt.Println(k)
	}

}

// Split function that splits by delimiters
func split(r rune) bool {
	return r == '\'' ||
		r == '"' ||
		r == ',' ||
		r == '.' ||
		r == '|' ||
		r == '-' ||
		r == '_' ||
		r == '(' ||
		r == ')' ||
		r == '=' ||
		r == ';' ||
		r == ':' ||
		r == '[' ||
		r == ']' ||
		r == '/' ||
		r == '\\' ||
		r == '<' ||
		r == '>' ||
		r == ' '
}
