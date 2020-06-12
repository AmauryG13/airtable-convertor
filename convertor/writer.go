package convertor

import (
	"log"
	"os"
	"strings"
)

// Writer is holding the writing logic
type Writer struct {
	Filepath string
	Options  Options
	file     *os.File
	channel  chan []string
}

// Options is holding writer options
type Options struct {
	EndOfLine string
	Separator string
}

// NewOptions is creating Options
func NewOptions(options map[string]string) *Options {
	return &Options{
		EndOfLine: options["eol"],
		Separator: options["sep"],
	}
}

// NewWriter is creating a new Writer
func NewWriter(filepath string, options map[string]string, cWrite chan []string) *Writer {
	wOptions := NewOptions(options)

	return &Writer{
		Filepath: filepath,
		Options:  *wOptions,
		channel:  cWrite,
	}
}

func (w *Writer) setFileID(fileID *os.File) {
	w.file = fileID
}

func (w *Writer) getFileID() *os.File {
	return w.file
}

// Open is creating a new file with the name selected
func (w *Writer) Open() {
	f, err := os.Create(w.Filepath)

	if err != nil {
		log.Fatal(err)
	}

	w.setFileID(f)
}

// Write is writing row to the new file
func (w *Writer) Write() {
	w.Open()
	defer w.Close()

	log.Println("writer: starting")

	for {
		log.Println("writer: started")

		select {
		case row, status := <-w.channel:
			log.Printf("writer: %s received %s \n", status, row)

			concatenateRow := strings.Join(row, w.Options.Separator)

			fullRow := concatenateRow + w.Options.EndOfLine
			log.Println("writer: joining", fullRow)

			_, err := w.file.WriteString(fullRow)
			log.Println("writer: writing")

			if err != nil {
				log.Fatal("writing error: ", err)
			}

			if !status {
				break
			}
		}
	}

}

// Close is closing the file opened identified by its file descriptor
func (w *Writer) Close() {
	err := w.getFileID().Close()

	if err != nil {
		log.Fatal(err)
	}
}
