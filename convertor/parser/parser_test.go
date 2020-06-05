package parser

import (
  "testing"
)

func TestParser(t *testing.T) {
  t.Run("Creating a new parser", func(t *testing.T) {
    parser := NewParser([]string{})

    chars := [][]string{}
    chars = append(chars, []string{"Hello; ", "World \n"})
    chars = append(chars, []string{"I'm", "Golang;"})

    content := parser.Parse(chars)

    t.Log(content)
  })
}
