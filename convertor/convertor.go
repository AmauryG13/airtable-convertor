package convertor

import (
	"path/filepath"

	"github.com/amauryg13/airtor/convertor/formatter"
	"github.com/amauryg13/airtor/convertor/parser"
	"github.com/amauryg13/airtor/convertor/reader"
	"github.com/amauryg13/airtor/convertor/writer"
)

// Convertor is the main struct
type Convertor struct {
	Input  string
	Output string
	Reader *reader.Reader
	Parser *parser.Parser
	Format *formatter.Formatter
	Writer *writer.Writer
}

// NewConvertor is creating a new convertor
func NewConvertor(filePath string, addChars []string, options map[string]string) *Convertor {
	dir, file := filepath.Split(filePath)
	ext := filepath.Ext(file)

	oFileName := file[:len(file)-len(ext)] + "_" + addOnOutToken + ext
	oFilePath := filepath.Join(dir, oFileName)

	return &Convertor{
		Input:  filePath,
		Output: oFilePath,
		Reader: reader.NewReader(filePath),
		Parser: parser.NewParser(addChars),
		Format: formatter.NewFormatter(ext, options),
		Writer: writer.NewWriter(oFilePath, options),
	}
}

// Run is running the conversion logic
func (c *Convertor) Run() {
	buffer := c.Reader.Read()

	content := c.Format.Read(buffer)

	updatedContent := c.Parser.Parse(content)

	for _, record := range updatedContent {
		c.Writer.Write(record)
	}
}
