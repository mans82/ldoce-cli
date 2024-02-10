package parser

import (
	"net/http"
	"testing"

	"github.com/mans82/ldoce-cli/internal/commons"
	"github.com/stretchr/testify/assert"
)

func TestParseEntry(t *testing.T) {

	t.Parallel()

	client := &http.Client{}

	tests := []struct {
		url        string
		subEntries []SubEntry
	}{
		{
			url: "https://www.ldoceonline.com/dictionary/bucket",
			subEntries: []SubEntry{
				{
					HyphenatedText: "buck‧et",
					IPA:            "/ˈbʌkɪt/",
					Type:           "noun",
					GrammerNotes:   "countable",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "bucket",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "",
					ExtraInfo:      "",
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/take",
			subEntries: []SubEntry{
				{
					HyphenatedText: "take",
					IPA:            "/teɪk/",
					Type:           "verb",
					GrammerNotes:   "",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "take",
					IPA:            "/teɪk/",
					Type:           "noun",
					GrammerNotes:   "countable usually singular",
					ExtraInfo:      "informal",
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/shadow",
			subEntries: []SubEntry{
				{
					HyphenatedText: "shad‧ow",
					IPA:            "/ˈʃædəʊ $ -doʊ/",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "shadow",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "shadow",
					IPA:            "",
					Type:           "adjective",
					GrammerNotes:   "only before noun",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "shad‧ow",
					IPA:            "/ˈʃædəʊ-doʊ/",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/triumph",
			subEntries: []SubEntry{
				{
					HyphenatedText: "triumph",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "intransitive",
					ExtraInfo:      "formal",
				},
				{
					HyphenatedText: "tri‧umph",
					IPA:            "/ˈtraɪəmf/",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
				},
				{
					HyphenatedText: "Triumph",
					IPA:            "",
					Type:           "",
					GrammerNotes:   "",
					ExtraInfo:      "trademark",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.url, func(t *testing.T) {

			t.Parallel()

			req, err := http.NewRequest("GET", test.url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}

			req.Header.Set("User-Agent", commons.UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request: %v", err)
			}

			entries, err := ParseEntry(resp.Body)
			if err != nil {
				t.Errorf("Error parsing entry: %v", err)
			}

			assert.Equal(t, test.subEntries, entries.SubEntries)

		})
	}
}
