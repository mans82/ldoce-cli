package parser

import "slices"

type QueryResult struct {
	Entries []Entry `yaml:"entries"`
}

func (q *QueryResult) FilterByTypes(entryTypes ...string) *QueryResult {
	filteredEntries := make([]Entry, 0)
	for _, entry := range q.Entries {
		if slices.Contains(entryTypes, entry.Type) {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return &QueryResult{Entries: filteredEntries}
}
