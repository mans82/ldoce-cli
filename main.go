package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mans82/ldoce-cli/pkg/formatter"
	"github.com/mans82/ldoce-cli/pkg/lookup"
)

type AppConfig struct {
	query        string
	entryTypes   []string
	showAllTypes bool
}

func main() {

	config := parseArgs()

	queryResult, err := lookup.LookupDefault(config.query)
	if err != nil {
		panic(err)
	}

	if !config.showAllTypes {
		queryResult = queryResult.FilterByTypes(config.entryTypes...)
	}

	bufferedStdout := bufio.NewWriter(os.Stdout)
	defer bufferedStdout.Flush()

	if len(queryResult.Entries) == 0 {
		fmt.Fprintf(bufferedStdout,
			"No entries found for '%s'. Check spelling or remove filters.\n", config.query)
		return
	}

	for _, entry := range queryResult.Entries {
		formatter.PrintFormattedEntry(bufferedStdout, entry)
	}
}

func usage() {
	fmt.Printf("Usage: %s [options] <word>\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(2)

}

func parseArgs() *AppConfig {

	config := &AppConfig{
		query:        "",
		entryTypes:   []string{},
		showAllTypes: true,
	}

	flag.Usage = usage
	helpFlag := flag.Bool("help", false, "show this help message")
	nounFlag := flag.Bool("noun", false, "show noun entries")
	verbFlag := flag.Bool("verb", false, "show verb entries")
	adjectiveFlag := flag.Bool("adj", false, "show adjective entries")
	adverbFlag := flag.Bool("adv", false, "show adverb entries")

	flag.Parse()

	if *helpFlag {
		usage()
		os.Exit(0)
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

	config.query = flag.Arg(0)

	typeFlags := map[*bool]string{
		nounFlag:      "noun",
		verbFlag:      "verb",
		adjectiveFlag: "adjective",
		adverbFlag:    "adverb",
	}

	config.entryTypes = make([]string, 0)

	for flagPtr, entryType := range typeFlags {
		if *flagPtr {
			config.entryTypes = append(config.entryTypes, entryType)
			config.showAllTypes = false
		}
	}

	return config

}
