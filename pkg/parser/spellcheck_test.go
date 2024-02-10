package parser

import (
	"net/http"
	"strings"
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
		{
			url: "https://www.ldoceonline.com/spellcheck/english/?q=gerr",
			suggestions: []string{
				"err",
				"gear",
				"germ",
				"gears",
				"genre",
				"germs",
				"-eer",
				"-er",
				"-ery",
				"Aero",
			},
		},
		{
			url: "https://www.ldoceonline.com/spellcheck/english/?q=uncle%20sam",
			suggestions: []string{
				"Uncle Sam",
				"Uncle Tom",
				"unclean",
				"unclear",
				"uncles",
				"cubicle farm",
				"uncle",
				"uncleaner",
				"unclearer",
				"Uncle Remus",
			},
		},
		{
			url: "https://www.ldoceonline.com/spellcheck/english/?q=aaa",
			suggestions: []string{
				"aha",
				"aka",
				"baa",
				"Sanaa",
				"-ana",
				"Baja",
				"Cana",
				"Dada",
				"Java", // I wish there were no "Java" in the list (Yes I hate Java)
				"Lada",
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

			assert.Equal(t, http.StatusOK, resp.StatusCode)

			result, err := ParseSpellcheck(resp.Body)
			if err != nil {
				t.Errorf("Error parsing spellcheck list: %v", err)
			}

			assert.Equal(t, len(test.suggestions), len(result.Suggestions))

			for i, suggestion := range result.Suggestions {
				urlSuffix := strings.Replace(test.suggestions[i], " ", "+", -1)

				assert.Equal(t, test.suggestions[i], suggestion.Text)
				assert.Equal(t, "https://www.ldoceonline.com/search/direct/?q="+urlSuffix, suggestion.Url)
			}
		})
	}

}
