package convertor

import (
	"fmt"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("Creating a new Reader", func(t *testing.T) {
		filepath := "../tests/data.csv"
		readChan := make(chan []string)

		reader := NewReader(filepath, readChan)

		if got, want := reader.Filepath, filepath; got != want {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}

	})

	t.Run("Handling a file to read", func(t *testing.T) {
		filepath := "../tests/data.csv"
		readChan := make(chan []string)

		reader := NewReader(filepath, readChan)
		reader.Open()
		reader.Close()
	})

	t.Run("Reading a file", func(t *testing.T) {
		filepath := "../tests/data.csv"
		readChan := make(chan []string)

		reader := NewReader(filepath, readChan)

		go reader.Read()

		eof := false
		for {

			for _, record := range <-readChan {
				if strings.Contains(record, "EOF") {
					eof = true
					break
				}

				fmt.Println(record)
			}

			if eof {
				break
			}
		}
	})
}
