package formatter

import (
	"io"

	"github.com/fatih/color"
	"github.com/mans82/ldoce-cli/pkg/parser"
)

var redBold = color.New(color.FgHiRed, color.Bold).FprintfFunc()
var white = color.New(color.FgWhite).FprintfFunc()
var whiteBold = color.New(color.FgWhite, color.Bold).FprintfFunc()
var green = color.New(color.FgHiGreen).FprintfFunc()
var magenta = color.New(color.FgHiMagenta).FprintfFunc()

func PrintFormattedEntry(writer io.Writer, entry parser.SubEntry) {
	redBold(writer, "  %s", entry.HyphenatedText)
	if entry.IPA != "" {
		white(writer, " %s", entry.IPA)
	}
	white(writer, " - ")
	if entry.Type != "" {
		green(writer, "%s ", entry.Type)
	}
	if entry.GrammerNotes != "" {
		white(writer, "[")
		green(writer, "%s", entry.GrammerNotes)
		white(writer, "] ")
	}
	magenta(writer, "%s\n", entry.ExtraInfo)

	for _, definition := range entry.Definitions {
		PrintFormattedDefinition(writer, definition)
	}
}

func PrintFormattedDefinition(writer io.Writer, definition string) {
	whiteBold(writer, "  + ")
	white(writer, "%s\n", definition)
}
