package formatter

import (
	"io"
	"reflect"
)

// Format is the interface to ensure all formatters work the same way
type Format interface {
	ReadContent(io.Reader) [][]string
	ExportContent(content [][]string, separator string) []string
}

// Options is the struct holding the options of the formatter
type Options struct {
	EndOfLine string
	Separator string
}

// Formatter is the struct holding the logic for formatting content
type Formatter struct {
	Type    string
	Options Options
	List    map[string]Format
}

func getAllFormatters() map[string]Format {
	list := make(map[string]Format)

	for fmt, fmtFunc := range availableFormatters {
		list[fmt] = (reflect.ValueOf(fmtFunc).Call([]reflect.Value{})[0].Interface()).(Format)
	}

	return list
}

// NewOptions is creating Options struct
func NewOptions(options map[string]string) *Options {
	return &Options{
		EndOfLine: options["eol"],
		Separator: options["sep"],
	}
}

// NewFormatter returns a new Formatter struct
func NewFormatter(formatType string, formatOptions map[string]string) *Formatter {
	newOptions := NewOptions(formatOptions)
	fmtList := getAllFormatters()

	return &Formatter{
		Type:    formatType,
		Options: *newOptions,
		List:    fmtList,
	}
}

func (formatter *Formatter) Read(fileReader io.Reader) [][]string {
	IFormat := formatter.List[formatter.Type]
	return IFormat.ReadContent(fileReader)
}

func (formatter *Formatter) Write(content [][]string) []string {
	IFormat := formatter.List[formatter.Type]
	return IFormat.ExportContent(content, formatter.Options.Separator)
}
