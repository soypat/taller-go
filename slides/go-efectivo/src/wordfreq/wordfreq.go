package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/bbalet/stopwords"
)

type Word struct {
	Count   int
	Palabra string
}

type Words []Word

func (a Words) Len() int           { return len(a) }
func (a Words) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Words) Less(i, j int) bool { return a[i].Count < a[j].Count }

func main() {
	fp, err := os.Open("covid.txt")
	if err != nil {
		panic(err)
	}
	buff, _ := ioutil.ReadAll(fp)
	text := string(buff) // convierto un slice de []byte a string
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, "\"", "")
	text = strings.ReplaceAll(text, "?", "")
	text = stopwords.CleanString(text, "en", false) // en for english

	palabras := strings.Fields(text) // str.split() de python

	wordcount := map[string]int{}
	for _, palabra := range palabras {
		_, presente := wordcount[palabra]
		if !presente {
			wordcount[palabra] = 0
		}
		wordcount[palabra]++
	}
	var wc Words
	for k, v := range wordcount {
		w := Word{Palabra: k,
			Count: v,
		}
		wc = append(wc, w)
	}

	sort.Sort(wc) // https://golang.org/pkg/sort/ (sort golang)
	for _, w := range wc {
		fmt.Printf("%v, %v\n", w.Count, w.Palabra)
	}

}
