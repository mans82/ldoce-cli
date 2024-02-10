package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Suggestion struct {
	Text string
	Url  string
}

type Spellcheck struct {
	Suggestions []Suggestion
}

func ParseSpellcheck(htmlTextReader io.Reader) (*Spellcheck, error) {

	doc, err := goquery.NewDocumentFromReader(htmlTextReader)
	if err != nil {
		return nil, fmt.Errorf("Error parsing html: %v", err)
	}

	suggestionSelection := doc.Find("ul.didyoumean > li > a")

	result := make([]Suggestion, 0, suggestionSelection.Length())

	suggestionSelection.Each(func(i int, s *goquery.Selection) {
		suggestionText := strings.TrimSpace(s.Text())
		suggestionUrl, _ := s.Attr("href")
		suggestionUrl = "https://www.ldoceonline.com" + suggestionUrl

		result = append(result, Suggestion{Text: suggestionText, Url: suggestionUrl})
	})

	return &Spellcheck{Suggestions: result}, nil
}
