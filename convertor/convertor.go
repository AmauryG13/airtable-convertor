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
	Base      string
	Dir       string
	Filename  string
	Extension string
}

func newInfo(base string, dir string, name string, ext string) *Info {
	return &Info{
		Base:      base,
		Dir:       dir,
		Filename:  name,
		Extension: ext,
	}
}

func newChannels() map[string]chan []string {
	Channels := make(map[string]chan []string)

	Channels["read"] = make(chan []string)
	Channels["write"] = make(chan []string)

	return Channels
}

func handlePath(dir string) (string, string) {
	if !filepath.IsAbs(dir) {
		base, err := os.Getwd()

		if err != nil {
			log.Fatal(err)
		}

		return base, dir
	}

	return dir, ""
}

// New is creating a new convertor
func New(filePath string, addChars []string, options map[string]string) *Convertor {
	path, file := filepath.Split(filePath)
	base, dir := handlePath(path)

	ext := filepath.Ext(file)

	cInfo := newInfo(base, dir, file, ext)

	iFilePath := filepath.Join(base, dir, file)

	oFileName := file[:len(file)-len(ext)] + "_" + addOnOutToken + ext
	oFilePath := filepath.Join(base, dir, oFileName)

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
	go c.writer.Write()
	go c.parser.Parse()
	c.reader.Read()
}
