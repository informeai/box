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
		title := doc.Find("h1", "id", "firstHeading")
		body := doc.Find("div", "id", "bodyContent")
		ps := body.FindAll("p")
		refs := body.Find("div", "class", "reflist").FindAll("a")
		re := regexp.MustCompile(`\[[0-9]*]`)
		fmt.Printf("%v\n\n", strings.ToUpper(title.Text()))
		for _, p := range ps {
			fmt.Println(re.ReplaceAllLiteralString(p.FullText(), ""))
		}
		fmt.Printf("REFERENCES:\n\n")
		for _, ref := range refs {
			if len(ref.Text()) > 4 {

				fmt.Printf("%v -> %v\n", ref.Text(), ref.Attrs()["href"])
			}
		}
	}
}
