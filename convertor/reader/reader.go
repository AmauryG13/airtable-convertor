package reader

import (
	"bufio"
	"io"
	"log"
	"os"
	"path"
)

//Reader struct
type Reader struct {
	Filepath string
	File     *os.File

	IsOpened bool

	filename string
}

func isExistingFile(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	return os.IsNotExist(err), err
}

func isReadableFile(filename string) (string, error) {
	cwd, _ := os.Getwd()
	filepath := path.Join(cwd, filename)

	status, err := isExistingFile(filepath)

	if status {
		return "", err
	}

	return filepath, err
}

func openInputFile(filepath string) *os.File {
	f, err := os.Open(filepath)

	if err != nil {
		log.Fatal("Cannot open file")
	}

	return f
}

//NewReader creates a new reader instance
func NewReader(filename string) *Reader {
	filepath, err := isReadableFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	fileID := openInputFile(filepath)

	return &Reader{
		Filepath: filepath,
		File:     fileID,
		IsOpened: true,
		filename: filename,
	}
}

func (r *Reader) Read() io.Reader {
	return bufio.NewReader(r.File)
}
