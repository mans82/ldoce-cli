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
						Definitions: []string{
							"to let someone have something as a present, or to provide something for someone",
							"to put something in someone’s hand",
							"to allow or make it possible for someone to do something",
							"to tell someone information or details about something, or to tell someone what they should do",
							"to do something by making a movement with your hand, face, body etc",
							"to make a speech, perform a piece of music etc for a group of people",
							"to make someone have a feeling",
							"to make someone have problems",
							"to infect someone with the same illness that you have",
							"to organize a social event such as a party",
							"to tell someone to do a job or piece of work",
							"to make someone or something have a particular quality",
							"to pay a particular amount of money for something",
							"to behave towards someone in a way that shows you have a particular attitude or feeling towards them",
							"to officially say that someone must have a particular punishment",
							"to be willing to change what you think or do according to what else happens",
							"to state what your official decision or judgment is, for example in a game",
							"to decide that someone should have a particular score or mark for something that they have done",
							"if a material gives, it bends or stretches when you put pressure on it",
							"if something gives, it breaks or moves away suddenly because of weight or pressure on it",
							"if a woman gives herself to a man, she has sex with him",
						},
					},
					{
						HyphenatedText: "give",
						IPA:            "",
						Type:           "noun",
						GrammerNotes:   "uncountable",
						ExtraInfo:      "",
						Definitions: []string{
							"the ability of a material or substance to bend or stretch when put under pressure",
						},
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
						Definitions: []string{
							"nerves are parts inside your body which look like threads and carry messages between the brain and other parts of the body",
							"courage and confidence in a dangerous, difficult, or frightening situation",
							"if you say someone has a nerve, you mean that they have done something unsuitable or impolite, without seeming to be embarrassed about behaving in this way",
						},
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
						Definitions: []string{
							"together with",
							"doing something with someone else as an equal or with less responsibility",
						},
					},
					{
						HyphenatedText: "Co.",
						IPA:            "/kəʊ $ koʊ/",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Definitions: []string{
							"the abbreviation of company",
							"the written abbreviation of county",
						},
					},
					{
						HyphenatedText: "CO",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Definitions: []string{
							"the written abbreviation of Colorado",
						},
					},
					{
						HyphenatedText: "C.O.",
						IPA:            "/ˌsiː ˈəʊ◂ $ -ˈoʊ◂/",
						Type:           "noun",
						GrammerNotes:   "countable",
						ExtraInfo:      "",
						Definitions: []string{
							"an officer who is in charge of a military unit",
						},
					},
					{
						HyphenatedText: "Co",
						IPA:            "",
						Type:           "",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Definitions: []string{
							"written abbreviation for Company",
						},
					},
					{
						HyphenatedText: "co",
						IPA:            "/kəʊkoʊ/",
						Type:           "prefix",
						GrammerNotes:   "",
						ExtraInfo:      "",
						Definitions: []string{
							"added to the front of a noun to show that someone does a job with someone else",
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
