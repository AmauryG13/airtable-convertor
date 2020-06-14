package convertor

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"sync"
)

//Reader struct
type Reader struct {
	Filepath string
	file     *os.File
	channel  chan []string
}

func isExistingFile(filepath string) (bool, error) {
	_, err := os.Stat(filepath)

	return !os.IsNotExist(err), err
}

//NewReader creates a new reader instance
func NewReader(filepath string, cRead chan []string) *Reader {
	status, err := isExistingFile(filepath)

	if status == false {
		log.Fatal("This file '", filepath, "' doesn't exist")
	}

	if err != nil {
		log.Fatal(err)
	}

	return &Reader{
		Filepath: filepath,
		channel:  cRead,
	}
}

func (r *Reader) setFileID(fileID *os.File) {
	r.file = fileID
}

func (r *Reader) getFileID() *os.File {
	return r.file
}

// Open func is opened a file
func (r *Reader) Open() {
	f, err := os.Open(r.Filepath)

	if err != nil {
		log.Fatal(err)
	}

	r.setFileID(f)
}

func (r *Reader) buffery() io.Reader {
	return bufio.NewReader(r.getFileID())
}

func (r *Reader) Read(wg *sync.WaitGroup) {
	r.Open()
	defer r.Close()

	buffer := r.buffery()
	csvReader := csv.NewReader(buffer)

	log.Println("reader: starting")
	for {
		log.Println("reader: started")

		row, err := csvReader.Read()
		log.Println("reader: read", row)

		if err == io.EOF {
			close(r.channel)
			log.Println("reader: close channel")

			wg.Wait()
			break
		}

		if err != nil {
			log.Fatal("reading error:", err)
		}

		r.channel <- row
		log.Println("reader: sent")
	}
}

// Close func is closing a opened file
func (r *Reader) Close() {
	err := r.getFileID().Close()

	if err != nil {
		log.Fatal(err)
	}
}
