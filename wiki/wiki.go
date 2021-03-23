package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

func main() {
	if err := verifyArgs(os.Args); err != nil {
		fmt.Println(err)
	} else {

		url := parseFlagToUrl()
		resp, err := getPage(url)
		if err != nil {
			fmt.Println(err)
		}
		resp, err = parseToHTML(resp)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp)
	}

}

func verifyArgs(a []string) error {
	if len(a) != 5 {
		return errors.New("error: all params not exited\nUse: wiki -l <language> -w <word-search>")
	}
	return nil
}
func parseFlagToUrl() string {
	lang := flag.String("l", "en", "language for search wikipedia")
	word := flag.String("w", "wikipedia", "word for search")
	flag.Parse()
	return string("https://" + *lang + ".wikipedia.org/wiki/" + *word)
}
func getPage(s string) (string, error) {
	r, err := soup.Get(s)
	if err != nil {
		return "", errors.New("error: Not get url")
	}
	return r, nil
}
func parseToHTML(s string) (string, error) {
	var text string
	doc := soup.HTMLParse(s)
	body := doc.Find("div", "id", "bodyContent")
	refs := body.Find("div", "class", "reflist")
	if refs.Error == nil {
		title := doc.Find("h1", "id", "firstHeading")
		text += fmt.Sprintf("%v:\n", strings.ToUpper(title.Text()))
		ps := body.FindAll("p")
		for _, p := range ps {
			t, err := regexTratament(p.FullText())
			if err != nil {
				return "", err
			}
			text += t
		}

		text += fmt.Sprintf("\n%v\n\n", strings.ToUpper("references:"))

		for _, ref := range refs.FindAll("a") {
			if len(ref.Text()) > 4 {

				text += fmt.Sprintf("%v -> %v\n", ref.Text(), ref.Attrs()["href"])

			}
		}
	} else {
		t := &text
		*t = "Not Found Word\nSorry :("
	}
	return text, nil
}

func regexTratament(s string) (string, error) {
	re, err := regexp.Compile(`\[[0-9]*]`)
	if err != nil {
		return "", errors.New("error: regex error compile")
	}
	return re.ReplaceAllLiteralString(s, ""), nil
}
