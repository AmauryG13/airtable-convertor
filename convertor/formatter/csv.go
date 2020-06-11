package formatter

import (
	"encoding/csv"
	"io"
	"log"
	"strings"
)

// CSVFormat is the struct holding the logic for csv formatting
type CSVFormat struct{}

// NewCSVFormat is creating a formatter
func NewCSVFormat() *CSVFormat {
	return &CSVFormat{}
}

// ReadContent is the function to read content from a CSV file
func (format *CSVFormat) ReadContent(fileReader io.Reader) [][]string {
	r := csv.NewReader(fileReader)

	records, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return records
}

// ExportContent is the function to convert the string array into a writable line array
func (format *CSVFormat) ExportContent(content [][]string, separator string) []string {
	var export []string

	for row := 0; row < len(content); row++ {
		export[row] = strings.Join(content[row][:], separator)
	}

	return export
}
