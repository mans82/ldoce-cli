package parser

import (
	"fmt"
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type SubEntry struct {
	HyphenatedText string
	IPA            string
	Type           string
	GrammerNotes   string
	ExtraInfo      string
	Definitions    []string
}

type Entry struct {
	SubEntries []SubEntry
}

func ParseEntry(htmlTextReader io.Reader) (*Entry, error) {

	doc, err := goquery.NewDocumentFromReader(htmlTextReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %v", err)
	}

	result := make([]SubEntry, 0)

	doc.Find(".Entry").Each(func(i int, s *goquery.Selection) {

		head := s.Find("span.Head")

		hyphenatedTextSelection := head.Find("span.HYPHENATION")

		if hyphenatedTextSelection.Length() == 0 {
			return
		}

		currentEntry := SubEntry{}

		currentEntry.HyphenatedText = hyphenatedTextSelection.Text()
		currentEntry.IPA = tryExtract(head, "span.Head > span.PronCodes")
		currentEntry.Type = tryExtract(head, "span.POS")
		currentEntry.GrammerNotes = strings.Trim(tryExtract(head, "span.GRAM"), "[]")
		currentEntry.ExtraInfo = tryExtract(head, "span.REGISTERLAB")

		s.Find(".Sense").Each(func(i int, node *goquery.Selection) {

			definition := strings.TrimSpace(node.Find(".DEF").Text())
			if definition == "" {
				return
			}

			definition += " " + strings.TrimSpace(node.Find(".DEF + .FULLFORM").Text())

			definition = strings.TrimSpace(definition)

			currentEntry.Definitions = append(currentEntry.Definitions, definition)
		})

		result = append(result, currentEntry)
	})

	return &Entry{SubEntries: result}, nil
}

func tryExtract(s *goquery.Selection, selector string) string {

	result := ""

	s.Find(selector).Each(func(i int, node *goquery.Selection) {
		result += node.Text()
	})

	return strings.TrimSpace(result)
}
