package parser

import (
	"fmt"
	"io"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	HyphenatedText string
	IPA            string
	Type           string
	GrammerNotes   string
	ExtraInfo      string
}

func ParseEntry(htmlTextReader io.Reader) ([]Entry, error) {

	doc, err := goquery.NewDocumentFromReader(htmlTextReader)
	if err != nil {
		return nil, fmt.Errorf("Error parsing html: %v", err)
	}

	result := make([]Entry, 0)

	doc.Find("span.Head").Each(func(i int, s *goquery.Selection) {

		hyphenatedTextSelection := s.Find("span.HYPHENATION")

		if len(hyphenatedTextSelection.Nodes) == 0 {
			return
		}

		currentEntry := Entry{}

		currentEntry.HyphenatedText = hyphenatedTextSelection.Text()
		currentEntry.IPA = tryExtract(s, "span.PronCodes")
		currentEntry.Type = tryExtract(s, "span.POS")
		currentEntry.GrammerNotes = tryExtract(s, "span.GRAM")
		currentEntry.ExtraInfo = tryExtract(s, "span.REGISTERLAB")

		result = append(result, currentEntry)
	})

	return result, nil
}

func tryExtract(s *goquery.Selection, selector string) string {

	result := ""

	s.Find(selector).Each(func(i int, node *goquery.Selection) {
		result += node.Text()
	})

	return result
}
