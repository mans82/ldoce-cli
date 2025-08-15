package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/mans82/ldoce-cli/pkg/formatter"
	"github.com/mans82/ldoce-cli/pkg/lookup"
)


func main() {

	parseArgs()

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

func usage() {
	fmt.Printf("Usage: %s [options] <word>\n\n", os.Args[0])
	fmt.Println("Options:")
	flag.PrintDefaults()
	os.Exit(2)

}

func parseArgs() {
	flag.Usage = usage
	helpFlag := flag.Bool("help", false, "show this help message")

	flag.Parse()

	if *helpFlag {
		usage()
	}

	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(2)
	}

}