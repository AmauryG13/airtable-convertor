package convertor

import (
	"strings"
)

// Parser is the struct holding logic
type Parser struct {
	ToBeRemoved  []string
	readChannel  chan []string
	writeChannel chan []string
}

func dispatchChannels(channels map[string]chan []string) (chan []string, chan []string) {
	rChan := channels["read"]
	wChan := channels["write"]

	return rChan, wChan
}

// NewParser is created a new Parser
func NewParser(removingChars []string, channels map[string]chan []string) *Parser {
	var remChars []string
	remChars = append(unWantedChars[:], removingChars...)

	rChan, wChan := dispatchChannels(channels)

	return &Parser{
		ToBeRemoved:  remChars,
		readChannel:  rChan,
		writeChannel: wChan,
	}
}

func removeChar(record string, toRemovedChars []string) string {
	for _, val := range toRemovedChars {
		record = strings.ReplaceAll(record, val, "")
	}

	return record
}

// Parse is the function used to parse content
func (p *Parser) Parse() {
	eof := false

	for {
		row := <-p.readChannel

		for idx, record := range row {
			if strings.Contains(record, "EOF") {
				eof = true
			}

			row[idx] = removeChar(record, p.ToBeRemoved)
		}

		p.writeChannel <- row

		if eof {
			break
		}
	}
}
