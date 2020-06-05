package reader

import (
  "testing"
  "path"
  "os"
)

func TestReader(t *testing.T) {
  t.Run("Reading a new existing file", func(t *testing.T) {
    filename := "reader.go"
    r := NewReader(filename)

    cwd,_ := os.Getwd()
    want := path.Join(cwd, filename)

    if got := r.Filepath; got != want {
      t.Errorf("GOT: %v; WANT: %v", got, want)
    }
  })
}
