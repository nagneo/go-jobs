package main

import (
	"fmt"

	"github.com/nagneo/go-jobs/mydict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil {
		fmt.Println(err)
	}

	err = dictionary.Delete(word)
	if err != nil {
		fmt.Println(err)
	}

	_, err = dictionary.Search(word)
	if err != nil {
		fmt.Println(err)
	}
}
