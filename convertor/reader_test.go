package convertor

import (
	"fmt"
	"sync"
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

		var wg sync.WaitGroup
		wg.Add(1)

		go reader.Read(&wg)

	loop:
		for {
			select {
			case row, status := <-readChan:
				if !status {
					wg.Done()
					break loop
				}

				for _, record := range row {
					fmt.Println(record)
				}
			}
		}
	})
}
