package output

import (
	"io"
	"os"

	"github.com/shellrausch/gofuzzy/fuzz/client"
)

var cli Writer
var outputWriter Writer
var outputFile io.Writer
var supportedFormats map[string]bool

func init() {
	supportedFormats = map[string]bool{"csv": true, "txt": true, "json": true}
}

// Writer must be implemented by every output format which wants to write.
type Writer interface {
	init()
	write(*client.Result)
	writeProgress(*client.Progress)
	close()
}

// Formats returns all available and supported output formats to which gofuzzy can write to.
func Formats() map[string]bool {
	return supportedFormats
}

// SetOutput sets the output file and decides on which output media
// the results should be shown. We always output on the CLI, also if another
// output media is provided.
func SetOutput(filename, outputFormat string) {
	outputFile, _ = os.Create(filename)

	var ow Writer

	switch outputFormat {
	case "csv":
		ow = csv{}
	case "txt":
		ow = txt{}
	case "json":
		ow = json{}
	default:
		ow = null{}
	}

	ow.init()
	outputWriter = ow

	// We want always a CLI output.
	cli = tabCli{}
	cli.init()
}

// Write just writes the result to the defined output and additionaly to the CLI.
func Write(r *client.Result) {
	outputWriter.write(r)
	cli.write(r)
}

// WriteProgress just writes the progress to the defined output and additionaly to the CLI.
func WriteProgress(pr *client.Progress) {
	outputWriter.writeProgress(pr)
	cli.writeProgress(pr)
}

// Close closes output writer.
func Close() {
	outputWriter.close()
}
