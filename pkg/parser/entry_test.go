package parser

import (
	"net/http"
	"testing"
)

const userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3"

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

			req.Header.Set("User-Agent", userAgent)

			resp, err := client.Do(req)
			if err != nil {
				t.Errorf("Error sending request: %v", err)
			}

			entries, err := ParseEntry(resp.Body)
			if err != nil {
				t.Errorf("Error parsing entry: %v", err)
			}

			if len(entries) != test.entryCount {
				t.Errorf("Expected %d entries, got %d", test.entryCount, len(entries))
			}
		})
	}
}
