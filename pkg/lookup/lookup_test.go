package lookup

import (
	"testing"

	"github.com/mans82/ldoce-cli/pkg/parser"
	"github.com/stretchr/testify/assert"
)

func TestLookUpDefault(t *testing.T) {
	t.Parallel()

	tests := []struct {
		word  string
		entry parser.Entry
	}{
		{
			word: "bucket",
			entry: parser.Entry{
				SubEntries: []parser.SubEntry{
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
		},
		{
			word: "give",
			entry: parser.Entry{
				SubEntries: []parser.SubEntry{
					{
						HyphenatedText: "give",
						IPA:            "/ɡɪv/",
						Type:           "verb",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "give",
						IPA:            "",
						Type:           "noun",
						GrammerNotes:   "uncountable",
						ExtraInfo:      "",
					},
				},
			},
		},
		{
			word: "nerv",
			entry: parser.Entry{
				SubEntries: []parser.SubEntry{
					{
						HyphenatedText: "nerve",
						IPA:            "/nɜːv $ nɜːrv/",
						Type:           "noun",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "nerve",
						IPA:            "",
						Type:           "verb",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
				},
			},
		},
		{
			word: "co",
			entry: parser.Entry{
				SubEntries: []parser.SubEntry{
					{
						HyphenatedText: "co-",
						IPA:            "/kəʊ $ koʊ/",
						Type:           "prefix",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "Co.",
						IPA:            "/kəʊ $ koʊ/",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "CO",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "C.O.",
						IPA:            "/ˌsiː ˈəʊ◂ $ -ˈoʊ◂/",
						Type:           "noun",
						GrammerNotes:   "countable",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "Co",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
					{
						HyphenatedText: "co",
						IPA:            "/kəʊkoʊ/",
						Type:           "prefix",
						GrammerNotes:   "",
						ExtraInfo:      "",
					},
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.word, func(t *testing.T) {
			t.Parallel()
			entry, err := LookupDefault(test.word)
			if err != nil {
				t.Errorf("Error looking up: %v", err)
			}

			assert.Equal(t, test.entry, *entry)
		})
	}
}
