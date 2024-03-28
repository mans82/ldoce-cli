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
					Definitions: []string{
						"an open container with a handle, used for carrying and holding things, especially liquids",
						"the quantity of liquid that a bucket can hold",
						"a part of a machine shaped like a large bucket and used for moving earth, water etc",
						"a large amount of something",
					},
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
					Definitions: []string{
						"to move or go with someone or something from one place to another",
						"used with a noun instead of using a verb to describe an action. For example, if you take a walk, you walk somewhere",
						"to remove something from a place",
						"if something takes a particular amount of time, money, effort etc, that amount of time etc is needed for it to happen or succeed",
						"to accept or choose something that is offered, suggested, or given to you",
						"to get hold of something in your hands",
						"to use a particular form of transport or a particular road in order to go somewhere",
						"to study a particular subject in school or college for an examination",
						"to do an examination or test",
						"to be the correct or suitable size, type etc for a particular person or thing",
						"to collect or gather something for a particular purpose",
						"to react to someone or something or consider them in a particular way",
						"to have or experience a particular feeling",
						"to get possession or control of something",
						"to swallow, breathe in, inject etc a drug or medicine",
						"to make someone or something go to a higher level or position",
						"to measure the amount, level, rate etc of something",
						"to make a number smaller by a particular amount",
						"if a shop, business etc takes a particular amount of money, it receives that amount of money from its customers",
						"to teach a particular group of students in a school or college",
						"to write down information",
						"if a man takes someone, he has sex with them",
						"if a treatment, dye, drug etc takes, it begins to work successfully",
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Definitions: []string{
						"an occasion when a film scene, song, action etc is recorded",
						"the amount of money earned by a shop or business in a particular period of time",
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "/teɪk/",
					Type:           "noun",
					GrammerNotes:   "countable usually singular",
					ExtraInfo:      "informal",
					Definitions: []string{
						"the amount of money earned by a business in a particular period of time",
						"to be willing to do something wrong or illegal in return for money",
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
					Definitions: []string{
						"to subtract one number from another number",
						"if a business takes or takes in a particular amount of money, it earns that money from selling its goods and services",
					},
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
					Definitions: []string{
						"the dark shape that someone or something makes on a surface when they are between that surface and the light",
						"darkness caused by something preventing light from reaching a place",
						"the bad effect or influence that something has, which makes other things seem less enjoyable, attractive, or impressive",
					},
				},
				{
					HyphenatedText: "shadow",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
					Definitions: []string{
						"to follow someone closely in order to watch what they are doing",
						"to cover something with a shadow, or make it dark",
					},
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
					Definitions: []string{
						"to watch someone very closely or work with them in order to learn how they do their job",
						"to change at the same rate or in the same way as something",
					},
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
					Definitions: []string{
						"to gain a victory or success after a difficult struggle",
					},
				},
				{
					HyphenatedText: "tri‧umph",
					IPA:            "/ˈtraɪəmf/",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Definitions: []string{
						"an important victory or success after a difficult struggle",
						"a feeling of pleasure and satisfaction that you get from victory or success",
						"a very successful example of something",
					},
				},
				{
					HyphenatedText: "Triumph",
					IPA:            "",
					Type:           "",
					GrammerNotes:   "",
					ExtraInfo:      "trademark",
					Definitions: []string{
						"a type of motorcycle made by the British company Triumph, which is known for being well-made in a traditional way. The Triumph company also used to make sports cars.",
					},
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
