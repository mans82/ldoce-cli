package formatter

import (
	"fmt"
	"io"
	"strings"

	"github.com/fatih/color"
	"github.com/mans82/ldoce-cli/pkg/parser"
)

var redBold = color.New(color.FgHiRed, color.Bold).FprintfFunc()
var white = color.New(color.FgWhite).FprintfFunc()
var whiteBold = color.New(color.FgWhite, color.Bold).FprintfFunc()
var green = color.New(color.FgHiGreen).FprintfFunc()
var magenta = color.New(color.FgHiMagenta).FprintfFunc()

func PrintFormattedEntry(writer io.Writer, entry parser.Entry) {
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

	for _, sense := range entry.Senses {
		PrintFormattedDefinition(writer, sense.Subsenses[0].Definition)
	}

	fmt.Fprintln(writer)
}

func PrintFormattedDefinition(writer io.Writer, definition string) {
	wrappedDefinition := wrapLines(definition, 80)

	whiteBold(writer, "  + ")
	white(writer, "%s\n", wrappedDefinition[0])
	for _, line := range wrappedDefinition[1:] {
		white(writer, "    %s\n", line)
	}
}

func wrapLines(s string, length int) []string {
	result := make([]string, 0)
	fields := strings.Fields(s)

	if len(fields) == 0 {
		return result
	}

	currentLine := fields[0]
	for _, field := range fields[1:] {
		if len(currentLine)+len(field) < length {
			currentLine += " " + field
		} else {
			result = append(result, currentLine)
			currentLine = field
		}
	}

	result = append(result, currentLine)
	return result
}
