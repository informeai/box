package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Use: wiki -l <language> -w <word-search>")
	} else {

		lang := flag.String("l", "en", "language for search wikipedia")
		word := flag.String("w", "wikipedia", "word for search")

		flag.Parse()
		url := string("https://" + *lang + ".wikipedia.org/wiki/" + *word)

		resp, err := http.Get(url)
		if err != nil {
			log.Fatalln("Error: ", err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
	}
}
