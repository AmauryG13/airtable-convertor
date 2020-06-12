package convertor

import (
	"log"
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

	log.Println("parser: starting")
	for {
		log.Println("parser: started")

		select {
		case row, status := <-p.readChannel:
			log.Printf("parser: %b received %s \n", status, row)

			for idx, record := range row {
				row[idx] = removeChar(record, p.ToBeRemoved)
			}

			log.Println("parser: parse", row)

			p.writeChannel <- row
			log.Println("parser: sent")

			if !status {
				break
			}

		}
	}
}
