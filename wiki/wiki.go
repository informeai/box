package wiki

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

type Wiki struct{
	Args []string
	Url string
}

func NewWiki(args []string) *Wiki{
	return &Wiki{Args:args}

}


func(w *Wiki) verifyArgs() error {
	if len(w.Args) != 5 {
		return errors.New("error: all params not exited\nUse: wiki -l <language> -w <word-search>")
	}
	return nil
}

func(w *Wiki) parseFlagToUrl() {
	lang := flag.String("l", "en", "language for search wikipedia")
	word := flag.String("w", "wikipedia", "word for search")
	flag.Parse()
	w.Url = fmt.Sprint("https://" + *lang + ".wikipedia.org/wiki/" + *word)
}

func(w *Wiki) GetPage() error {
	err := w.verifyArgs()
	if err != nil{
		return err
	}
	w.parseFlagToUrl()
	r, err := soup.Get(w.Url)
	if err != nil {
		return errors.New("error: Not get url")
	}
	resp , err := w.parseToHTML(r)
	if err != nil{
		return errors.New("error to parse html")
	}
	fmt.Println(resp)
	return nil
	
}
func(w *Wiki) parseToHTML(s string) (string, error) {
	var text string
	doc := soup.HTMLParse(s)
	body := doc.Find("div", "id", "bodyContent")
	refs := body.Find("div", "class", "reflist")
	if refs.Error == nil {
		title := doc.Find("h1", "id", "firstHeading")
		text += fmt.Sprintf("%v:\n", strings.ToUpper(title.Text()))
		ps := body.FindAll("p")
		for _, p := range ps {
			t, err := w.regexTratament(p.FullText())
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

func(w *Wiki) regexTratament(s string) (string, error) {
	re, err := regexp.Compile(`\[[0-9]*]`)
	if err != nil {
		return "", errors.New("error: regex error compile")
	}
	return re.ReplaceAllLiteralString(s, ""), nil
}
