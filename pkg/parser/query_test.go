package parser

import "testing"

func TestQueryResultFilterByType(t *testing.T) {
	t.Parallel()

	queryResult := &QueryResult{
		Entries: []Entry{
			{Type: "noun"},
			{Type: "verb"},
			{Type: "noun"},
		},
	}

	filteredResult := queryResult.FilterByTypes("noun")
	if len(filteredResult.Entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(filteredResult.Entries))
	}

	for _, entry := range filteredResult.Entries {
		if entry.Type != "noun" {
			t.Errorf("Expected entry type 'noun', got '%s'", entry.Type)
		}
	}

	emptyResult := queryResult.FilterByTypes("adjective")
	if len(emptyResult.Entries) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(emptyResult.Entries))
	}
}
