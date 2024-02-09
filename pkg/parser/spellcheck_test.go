package parser

import (
	"net/http"
	"testing"

	"github.com/mans82/ldoce-cli/internal/commons"
	"github.com/stretchr/testify/assert"
)

func TestParseSpellcheckList(t *testing.T) {

	t.Parallel()

	client := &http.Client{}

	tests := []struct {
		url         string
		suggestions []string
	}{
		{
			url: "https://www.ldoceonline.com/spellcheck/english/?q=berr",
			suggestions: []string{
				"berry",
				"bear",
				"beer",
				"berk",
				"burr",
				"err",
				"beard",
				"bears",
				"beers",
				"beery",
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

			result, err := ParseSpellcheck(resp.Body)
			if err != nil {
				t.Errorf("Error parsing spellcheck list: %v", err)
			}

			assert.Equal(t, len(test.suggestions), len(result.Suggestions))

			for i, suggestion := range result.Suggestions {
				assert.Equal(t, test.suggestions[i], suggestion.Text)
				assert.Equal(t, "https://ldoceonline.com/search/direct/?q="+test.suggestions[i], suggestion.Url)
			}
		})
	}

}
