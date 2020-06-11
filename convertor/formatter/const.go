package formatter

//
var availableFormatters = map[string]interface{}{
	"csv": NewCSVFormat,
}
