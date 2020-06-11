package writer

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"
)

// Writer is the struct holding all func
type Writer struct {
	Filepath string
	File     io.Writer
	Options  Options
}

// Options is
type Options struct {
	EndOfLine string
	Separator string
}

func createFile(filename string) (*io.Writer, error) {
	cwd, _ := os.Getwd()
	filepath := path.Join(cwd, filename)

	return os.Create(filepath)
}

// NewOptions is creating Options struct
func NewOptions(options map[string]string) *Options {
	return &Options{
		EndOfLine: options["eol"],
		Separator: options["sep"],
	}
}

// NewWriter is
func NewWriter(filepath string, options map[string]string) *Writer {
	fileID, err := createFile(filepath)

	if err != nil {
		log.Fatal(err)
	}

	wOptions := NewOptions(options)

	return &Writer{
		Filepath: filepath,
		File:     *fileID,
		Options:  *wOptions,
	}
}

// Write
func (w *Writer) Write(record []string) error {
	csvWriter := csv.NewWriter(w.File)

	if err := csvWriter.Write(record); err != nil {
		log.Fatal(err)
	}

	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatal(err)
	}
}
