package parser

import (
  "strings"
)

type Parser struct {
  ToBeRemoved []string
}

func NewParser(removingChars []string) (*Parser) {
  var remChars []string
  remChars = append(UnwantedChars[:], removingChars...)

  return &Parser{
    ToBeRemoved: remChars,
  }
}

func removeChar(record string, toRemovedChars []string) (string) {
  for _, val := range toRemovedChars {
    record = strings.ReplaceAll(record, val, "")
  }

  return record
}

// Parse is the function used to parse content
func (p *Parser) Parse(content [][]string) ([][]string) {
  for rIdx := range content {
    for cIdx := range content {
      content[rIdx][cIdx] = removeChar(content[rIdx][cIdx], p.ToBeRemoved)
    }
  }

  return content
}
