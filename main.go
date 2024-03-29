package main

import (
	"bufio"
	"os"

	"github.com/mans82/ldoce-cli/pkg/formatter"
	"github.com/mans82/ldoce-cli/pkg/lookup"
)

func main() {

	entries, err := lookup.LookupDefault(os.Args[1])

	if err != nil {
		panic(err)
	}

	bufferedStdout := bufio.NewWriter(os.Stdout)
	defer bufferedStdout.Flush()

	for _, entry := range entries.Entries {
		formatter.PrintFormattedEntry(bufferedStdout, entry)
	}
}
