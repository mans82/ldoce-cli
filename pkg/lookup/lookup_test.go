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
		entry parser.QueryResult
	}{
		{
			word: "bucket",
			entry: parser.QueryResult{
				Entries: []parser.Entry{
					{
						HyphenatedText: "buck‧et",
						IPA:            "/ˈbʌkɪt/",
						Type:           "noun",
						GrammerNotes:   "countable",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "an open container with a handle, used for carrying and holding things, especially liquids"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "the quantity of liquid that a bucket can hold"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "a part of a machine shaped like a large bucket and used for moving earth, water etc"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "a large amount of something"},
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
						Senses:         []parser.Sense{},
					},
				},
			},
		},
		{
			word: "give",
			entry: parser.QueryResult{
				Entries: []parser.Entry{
					{
						HyphenatedText: "give",
						IPA:            "/ɡɪv/",
						Type:           "verb",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to let someone have something as a present, or to provide something for someone"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to put something in someone’s hand"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to allow or make it possible for someone to do something"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to tell someone information or details about something, or to tell someone what they should do"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to do something by making a movement with your hand, face, body etc"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to make a speech, perform a piece of music etc for a group of people"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to make someone have a feeling"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to make someone have problems"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to infect someone with the same illness that you have"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to organize a social event such as a party"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to tell someone to do a job or piece of work"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to make someone or something have a particular quality"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to pay a particular amount of money for something"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to behave towards someone in a way that shows you have a particular attitude or feeling towards them"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to officially say that someone must have a particular punishment"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to be willing to change what you think or do according to what else happens"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to state what your official decision or judgment is, for example in a game"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "to decide that someone should have a particular score or mark for something that they have done"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "if a material gives, it bends or stretches when you put pressure on it"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "if something gives, it breaks or moves away suddenly because of weight or pressure on it"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "if a woman gives herself to a man, she has sex with him"},
								},
							},
						},
					},
					{
						HyphenatedText: "give",
						IPA:            "",
						Type:           "noun",
						GrammerNotes:   "uncountable",
						ExtraInfo:      "",
						Senses: []parser.Sense{{
							SignPost: "",
							Subsenses: []parser.SubSense{
								{Definition: "the ability of a material or substance to bend or stretch when put under pressure"},
							},
						},
						},
					},
				},
			},
		},
		{
			word: "nerv",
			entry: parser.QueryResult{
				Entries: []parser.Entry{
					{
						HyphenatedText: "nerve",
						IPA:            "/nɜːv $ nɜːrv/",
						Type:           "noun",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "nerves are parts inside your body which look like threads and carry messages between the brain and other parts of the body"},
								},
							}, {
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "courage and confidence in a dangerous, difficult, or frightening situation"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "if you say someone has a nerve, you mean that they have done something unsuitable or impolite, without seeming to be embarrassed about behaving in this way"},
								},
							},
						},
					},
					{
						HyphenatedText: "nerve",
						IPA:            "",
						Type:           "verb",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses:         []parser.Sense{},
					},
				},
			},
		},
		{
			word: "co",
			entry: parser.QueryResult{
				Entries: []parser.Entry{
					{
						HyphenatedText: "co-",
						IPA:            "/kəʊ $ koʊ/",
						Type:           "prefix",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "together with"},
								},
							}, {
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "doing something with someone else as an equal or with less responsibility"},
								},
							},
						},
					},
					{
						HyphenatedText: "Co.",
						IPA:            "/kəʊ $ koʊ/",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "the abbreviation of company"},
								},
							},
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "the written abbreviation of county"},
								},
							},
						},
					},
					{
						HyphenatedText: "CO",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "the written abbreviation of Colorado"},
								},
							},
						},
					},
					{
						HyphenatedText: "C.O.",
						IPA:            "/ˌsiː ˈəʊ◂ $ -ˈoʊ◂/",
						Type:           "noun",
						GrammerNotes:   "countable",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "an officer who is in charge of a military unit"},
								},
							},
						},
					},
					{
						HyphenatedText: "Co",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "written abbreviation for Company"},
								},
							},
						},
					},
					{
						HyphenatedText: "co",
						IPA:            "/kəʊkoʊ/",
						Type:           "prefix",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Senses: []parser.Sense{
							{
								SignPost: "",
								Subsenses: []parser.SubSense{
									{Definition: "added to the front of a noun to show that someone does a job with someone else"},
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
