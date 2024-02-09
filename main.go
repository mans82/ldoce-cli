package main

import (
	"os"

	"github.com/mans82/ldoce-cli/pkg/formatter"
	"github.com/mans82/ldoce-cli/pkg/lookup"
)

func main() {

	entries, err := lookup.LookupDefault(os.Args[1])

	if err != nil {
		panic(err)
	}

	for _, entry := range entries.SubEntries {
		formatter.PrintFormattedEntry(os.Stdout, entry)
	}
}
