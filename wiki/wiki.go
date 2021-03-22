package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Use: wiki -l <language> -w <word-search>")
	} else {

		lang := flag.String("l", "en", "language for search wikipedia")
		word := flag.String("w", "wikipedia", "word for search")

		flag.Parse()
		url := string("https://" + *lang + ".wikipedia.org/wiki/" + *word)

		resp, err := soup.Get(url)
		if err != nil {
			log.Fatalln("Error: Not get url")
		}
		doc := soup.HTMLParse(resp)
		body := doc.Find("div", "id", "bodyContent")
		re, err := regexp.Compile(`\[[0-9]*]`)
		if err != nil {
			log.Fatalln("ERROR: Regex error compile")
		}
		refs := body.Find("div", "class", "reflist")
		if refs.Error == nil {
			title := doc.Find("h1", "id", "firstHeading")
			fmt.Printf("%v\n\n", strings.ToUpper(title.Text()))
			ps := body.FindAll("p")
			for _, p := range ps {
				fmt.Println(re.ReplaceAllLiteralString(p.FullText(), ""))
			}
			fmt.Printf("REFERENCES:\n\n")

			for _, ref := range refs.FindAll("a") {
				if len(ref.Text()) > 4 {

					fmt.Printf("%v -> %v\n", ref.Text(), ref.Attrs()["href"])
				}
			}
		} else {
			fmt.Printf("Not Found Word\nSorry :(\n")
		}
	}
}
