package lookup

import (
	"os"
	"testing"

	"github.com/mans82/ldoce-cli/pkg/parser"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestLookUpDefault(t *testing.T) {

	t.Parallel()

	allEntries, err := parser.GetAllTestEntries("../parser/testdata/all-entries.yaml") // FIXME: Find a better way to organize test data

	if err != nil {
		t.Errorf("error getting test entries: %v", err)
	}

	testsYamlFile, err := os.ReadFile("./testdata/tests.yaml")
	if err != nil {
		t.Errorf("error reading test data: %v", err)
	}

	var tests []struct {
		Query        string `yaml:"query"`
		ExpectedWord string `yaml:"expectedWord"`
	} = nil

	err = yaml.Unmarshal(testsYamlFile, &tests)

	if err != nil {
		t.Errorf("error unmarshalling test data: %v", err)
	}

	for _, test := range tests {
		test := test
		t.Run(test.Query, func(t *testing.T) {
			t.Parallel()
			queryResult, err := LookupDefault(test.Query)
			if err != nil {
				t.Errorf("Error looking up: %v", err)
			}

			assert.Equal(t, true, queryResult.SpellingIsCorrect)
			assert.Equal(t, allEntries[test.ExpectedWord], queryResult.Entries)
			assert.Equal(t, 0, len(queryResult.Spellings))
		})
	}
}
