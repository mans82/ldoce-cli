package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mans82/ldoce-cli/pkg/formatter"
	"github.com/mans82/ldoce-cli/pkg/lookup"
	"github.com/mans82/ldoce-cli/pkg/parser"
)

type AppConfig struct {
	query         string
	entryTypes    []string
	showAllTypes  bool
	checkSpelling bool
}

func main() {

	config := parseArgs()

	var queryResult *parser.QueryResult
	var err error

	if config.checkSpelling {
		queryResult, err = lookup.Lookup(config.query)
	} else {
		queryResult, err = lookup.LookupDefault(config.query)
	}

	if err != nil {
		panic(err)
	}

	stdioWriter := os.Stdout

	if config.checkSpelling && !queryResult.SpellingIsCorrect {
		formatter.PrintSpellingSuggestions(stdioWriter, queryResult.Spellings)
		os.Exit(1)
	}

	if !config.showAllTypes {
		queryResult = queryResult.FilterByTypes(config.entryTypes...)
	}

	if len(queryResult.Entries) == 0 {
		fmt.Fprintf(stdioWriter,
			"No entries found for '%s'. Check spelling or remove filters.\n", config.query)
		return
	}

	for _, entry := range queryResult.Entries {
		formatter.PrintFormattedEntry(stdioWriter, entry)
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
	checkFlag := flag.Bool("check", false, "check spelling first")

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

	config.checkSpelling = *checkFlag

	return config

}
