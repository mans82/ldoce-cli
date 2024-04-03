package parser

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v3"
)

type SubSense struct {
	Definition string `yaml:"definition"`
}

type Sense struct {
	SignPost  string     `yaml:"signpost"`
	Subsenses []SubSense `yaml:"subsenses"`
}

type Entry struct {
	HyphenatedText string  `yaml:"hyphenatedText"`
	IPA            string  `yaml:"IPA"`
	Type           string  `yaml:"type"`
	GrammerNotes   string  `yaml:"grammerNotes"`
	ExtraInfo      string  `yaml:"extraInfo"`
	Senses         []Sense `yaml:"senses"`
}

type QueryResult struct {
	Entries []Entry
}

func ParseEntry(htmlTextReader io.Reader) (*QueryResult, error) {

	doc, err := goquery.NewDocumentFromReader(htmlTextReader)
	if err != nil {
		return nil, fmt.Errorf("error parsing html: %v", err)
	}

	entries := make([]Entry, 0)

	doc.Find(".Entry").Each(func(i int, s *goquery.Selection) {

		head := s.Find("span.Head")

		hyphenatedTextSelection := head.Find("span.HYPHENATION")

		if hyphenatedTextSelection.Length() == 0 {
			return
		}

		currentEntry := Entry{}

		currentEntry.HyphenatedText = hyphenatedTextSelection.Text()
		currentEntry.IPA = tryExtract(head, "span.Head > span.PronCodes")
		currentEntry.Type = tryExtract(head, "span.POS")
		currentEntry.GrammerNotes = strings.Trim(tryExtract(head, "span.GRAM"), "[]")
		currentEntry.ExtraInfo = tryExtract(head, "span.REGISTERLAB")

		senses := make([]Sense, 0)

		s.Find(".Sense").Each(func(i int, node *goquery.Selection) {

			var sense Sense
			subSensesNodes := node.Find(".Subsense")

			if subSensesNodes.Length() == 0 {

				definition := strings.TrimSpace(node.Find(".DEF").Text())
				if definition == "" {
					return
				}

				definition += " " + strings.TrimSpace(node.Find(".DEF + .FULLFORM").Text())

				definition = strings.TrimSpace(definition)

				sense = Sense{
					SignPost:  "",
					Subsenses: []SubSense{{Definition: definition}},
				}
			} else {
				subSenses := make([]SubSense, 0)

				subSensesNodes.Each(func(i int, node *goquery.Selection) {
					definition := strings.TrimSpace(node.Find(".DEF").Text())
					if definition == "" {
						return
					}

					definition += " " + strings.TrimSpace(node.Find(".DEF + .FULLFORM").Text())

					definition = strings.TrimSpace(definition)

					subSenses = append(subSenses, SubSense{Definition: definition})
				})

				sense = Sense{
					SignPost:  "",
					Subsenses: subSenses,
				}
			}

			senses = append(senses, sense)
		})

		currentEntry.Senses = senses

		entries = append(entries, currentEntry)
	})

	return &QueryResult{Entries: entries}, nil
}

func tryExtract(s *goquery.Selection, selector string) string {

	result := ""

	s.Find(selector).Each(func(i int, node *goquery.Selection) {
		result += node.Text()
	})

	return strings.TrimSpace(result)
}

func GetAllTestEntries(testcasesFilePath string) (map[string]Entry, error) {

	var allEntries map[string]Entry

	testYamlFile, err := os.ReadFile(testcasesFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading test data: %v", err)
	}

	err = yaml.Unmarshal(testYamlFile, &allEntries)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling test data: %v", err)
	}

	return allEntries, nil
}
