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
		entryCount int
	}{
		{
			url:        "https://www.ldoceonline.com/dictionary/bucket",
			entryCount: 2,
		},
		{
			url:        "https://www.ldoceonline.com/dictionary/take",
			entryCount: 4,
		},
		{
			url:        "https://www.ldoceonline.com/dictionary/shadow",
			entryCount: 4,
		},
		{
			url:        "https://www.ldoceonline.com/dictionary/triumph",
			entryCount: 3,
		},
		{
			url:        "https://www.ldoceonline.com/dictionary/bayonet",
			entryCount: 2,
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

			assert.Equal(t, test.entryCount, len(entries.SubEntries))

		})
	}
}
