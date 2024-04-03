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

	tests, err := GetAllTestEntries()

	if err != nil {
		t.Errorf("Error getting test entries: %v", err)
	}

	for word, test := range tests {
		test := test
		test_url := "https://www.ldoceonline.com/dictionary/" + word

		t.Run(word, func(t *testing.T) {

			t.Parallel()

			req, err := http.NewRequest("GET", test_url, nil)
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

			assert.Equal(t, test.SubEntries, entries.SubEntries)

		})
	}
}
