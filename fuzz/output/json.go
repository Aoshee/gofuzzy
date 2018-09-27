package output

import (
	// Avoids naming collision with our struct name. Yes it's dirty.
	jsonn "encoding/json"
	"fmt"
	"log"

	"github.com/shellrausch/gofuzzy/fuzz/client"
)

type json struct{}

var jsonResults []*client.Result

func (json) write(r *client.Result) {
	jsonResults = append(jsonResults, r)
}

func (json) close() {
	json, err := jsonn.Marshal(jsonResults)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(outputFile, string(json))
}

func (json) init()                            {}
func (json) writeProgress(p *client.Progress) {}
