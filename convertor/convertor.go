package convertor

import (
	"log"
	"os"
	"path/filepath"
)

// Convertor is the main struct
type Convertor struct {
	Input    string
	Output   string
	Info     *Info
	Channels map[string]chan []string

	reader *Reader
	parser *Parser
	writer *Writer
}

// Info is the struct providing some info on path
type Info struct {
	Dir       string
	Filename  string
	Extension string
}

func newInfo(dir string, name string, ext string) *Info {
	return &Info{
		Dir:       dir,
		Filename:  name,
		Extension: ext,
	}
}

func newChannels() map[string]chan []string {
	var Channels map[string]chan []string

	Channels["read"] = make(chan []string)
	Channels["write"] = make(chan []string)

	return Channels
}

// NewConvertor is creating a new convertor
func NewConvertor(filePath string, addChars []string, options map[string]string) *Convertor {
	dir, file := filepath.Split(filePath)

	if dir == "" {
		cwd, cwdErr := os.Getwd()
		dir = cwd

		if cwdErr == nil {
			log.Fatal(cwdErr)
		}
	}

	ext := filepath.Ext(file)

	cInfo := newInfo(dir, file, ext)

	iFilePath := filepath.Join(dir, file)

	oFileName := file[:len(file)-len(ext)] + "_" + addOnOutToken + ext
	oFilePath := filepath.Join(dir, oFileName)

	convertorChan := newChannels()

	return &Convertor{
		Input:    iFilePath,
		Output:   oFilePath,
		Info:     cInfo,
		Channels: convertorChan,
		reader:   NewReader(iFilePath, convertorChan["read"]),
		parser:   NewParser(addChars, convertorChan),
		writer:   NewWriter(oFilePath, options, convertorChan["write"]),
	}
}

// Run is running conversion logic
func (c *Convertor) Run() {
	go c.reader.Read()
	go c.parser.Parse()
	go c.writer.Write()
}
