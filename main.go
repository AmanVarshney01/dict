package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Definition struct {
	Definition string `json:"definition"`
	Example    string `json:"example"`
}

type Meaning struct {
	PartsOfSpeech string       `json:"partOfSpeech"`
	Definitions   []Definition `json:"definitions"`
	Synonyms      []string     `json:"synonyms"`
	Antonyms      []string     `json:"antonyms"`
}

type Word struct {
	Word     string    `json:"word"`
	Meanings []Meaning `json:"meanings"`
}

func main() {

	input := "hello"

	if len(os.Args) > 1 {
		input = os.Args[1]
	}

	res, err := http.Get("https://api.dictionaryapi.dev/api/v2/entries/en/" + input)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		panic("Api not working")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var word []Word
	err = json.Unmarshal(body, &word)
	if err != nil {
		panic(err)
	}

	for _, w := range word {
		fmt.Println("Word: ", w.Word)
		for _, m := range w.Meanings {
			fmt.Println("Part of Speech: ", m.PartsOfSpeech)
			for _, d := range m.Definitions {
				fmt.Println("Definition: ", d.Definition)
				fmt.Println("Example: ", d.Example)
			}
			fmt.Println("Synonyms: ", m.Synonyms)
			fmt.Println("Antonyms: ", m.Antonyms)
		}
	}

}
