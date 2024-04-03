package parser

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/mans82/ldoce-cli/internal/commons"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestParseSpellcheckList(t *testing.T) {

	t.Parallel()

	client := &http.Client{}

	var tests map[string][]string = nil
	testYamlFile, err := os.ReadFile("./testdata/spellcheck.yaml")
	if err != nil {
		t.Errorf("Error reading test data: %v", err)
	}

	err = yaml.Unmarshal(testYamlFile, &tests)
	if err != nil {
		t.Errorf("Error unmarshalling test data: %v", err)
	}

	for query, suggestions := range tests {

		suggestions := suggestions
		query = strings.ReplaceAll(query, "-", "+")
		url := "https://www.ldoceonline.com/spellcheck/english/?q=" + query

		t.Run(query, func(t *testing.T) {

			t.Parallel()

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				t.Errorf("Error creating request: %v", err)
			}

			req.Header.Set("User-Agent", commons.UserAgent)

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request: %v", err)
			}

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			result, err := ParseSpellcheck(resp.Body)
			if err != nil {
				t.Errorf("Error parsing spellcheck list: %v", err)
			}

			assert.Equal(t, len(suggestions), len(result.Suggestions))

			for i, suggestion := range result.Suggestions {

				urlSuffix := strings.Replace(suggestions[i], " ", "+", -1)

				assert.Equal(t, suggestions[i], suggestion.Text)
				assert.Equal(t, "https://www.ldoceonline.com/search/direct/?q="+urlSuffix, suggestion.Url)
			}
		})
	}

}
