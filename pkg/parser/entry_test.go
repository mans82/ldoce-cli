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
		url     string
		entries []Entry
	}{
		{
			url: "https://www.ldoceonline.com/dictionary/bucket",
			entries: []Entry{
				{
					HyphenatedText: "buck‧et",
					IPA:            "/ˈbʌkɪt/",
					Type:           "noun",
					GrammerNotes:   "countable",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "an open container with a handle, used for carrying and holding things, especially liquids",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "the quantity of liquid that a bucket can hold",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "a part of a machine shaped like a large bucket and used for moving earth, water etc",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "a large amount of something",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "bucket",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Senses:         []Sense{},
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/take",
			entries: []Entry{
				{
					HyphenatedText: "take",
					IPA:            "/teɪk/",
					Type:           "verb",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to move or go with someone or something from one place to another",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "used with a noun instead of using a verb to describe an action. For example, if you take a walk, you walk somewhere",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to remove something from a place",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "if something takes a particular amount of time, money, effort etc, that amount of time etc is needed for it to happen or succeed",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to accept or choose something that is offered, suggested, or given to you",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to get hold of something in your hands",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to use a particular form of transport or a particular road in order to go somewhere",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to study a particular subject in school or college for an examination",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to do an examination or test",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to be the correct or suitable size, type etc for a particular person or thing",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to collect or gather something for a particular purpose",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to react to someone or something or consider them in a particular way",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to have or experience a particular feeling",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to get possession or control of something",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to swallow, breathe in, inject etc a drug or medicine",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to make someone or something go to a higher level or position",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to measure the amount, level, rate etc of something",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to make a number smaller by a particular amount",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "if a shop, business etc takes a particular amount of money, it receives that amount of money from its customers",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to teach a particular group of students in a school or college",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to write down information",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "if a man takes someone, he has sex with them",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "if a treatment, dye, drug etc takes, it begins to work successfully",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "an occasion when a film scene, song, action etc is recorded",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "the amount of money earned by a shop or business in a particular period of time",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "/teɪk/",
					Type:           "noun",
					GrammerNotes:   "countable usually singular",
					ExtraInfo:      "informal",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "the amount of money earned by a business in a particular period of time",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to be willing to do something wrong or illegal in return for money",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "take",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to subtract one number from another number",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "if a business takes or takes in a particular amount of money, it earns that money from selling its goods and services",
								},
							},
						},
					},
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/shadow",
			entries: []Entry{
				{
					HyphenatedText: "shad‧ow",
					IPA:            "/ˈʃædəʊ $ -doʊ/",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "the dark shape that someone or something makes on a surface when they are between that surface and the light",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "darkness caused by something preventing light from reaching a place",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "the bad effect or influence that something has, which makes other things seem less enjoyable, attractive, or impressive",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "shadow",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to follow someone closely in order to watch what they are doing",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to cover something with a shadow, or make it dark",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "shadow",
					IPA:            "",
					Type:           "adjective",
					GrammerNotes:   "only before noun",
					ExtraInfo:      "",
					Senses:         []Sense{},
				},
				{
					HyphenatedText: "shad‧ow",
					IPA:            "/ˈʃædəʊ-doʊ/",
					Type:           "verb",
					GrammerNotes:   "transitive",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to watch someone very closely or work with them in order to learn how they do their job",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to change at the same rate or in the same way as something",
								},
							},
						},
					},
				},
			},
		},
		{
			url: "https://www.ldoceonline.com/dictionary/triumph",
			entries: []Entry{
				{
					HyphenatedText: "triumph",
					IPA:            "",
					Type:           "verb",
					GrammerNotes:   "intransitive",
					ExtraInfo:      "formal",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "to gain a victory or success after a difficult struggle",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "tri‧umph",
					IPA:            "/ˈtraɪəmf/",
					Type:           "noun",
					GrammerNotes:   "",
					ExtraInfo:      "",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "an important victory or success after a difficult struggle",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "a feeling of pleasure and satisfaction that you get from victory or success",
								},
							},
						},
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "a very successful example of something",
								},
							},
						},
					},
				},
				{
					HyphenatedText: "Triumph",
					IPA:            "",
					Type:           "",
					GrammerNotes:   "",
					ExtraInfo:      "trademark",
					Senses: []Sense{
						{
							SignPost: "",
							Subsenses: []SubSense{
								{
									Definition: "a type of motorcycle made by the British company Triumph, which is known for being well-made in a traditional way. The Triumph company also used to make sports cars.",
								},
							},
						},
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

			assert.Equal(t, test.entries, entries.Entries)

		})
	}
}
