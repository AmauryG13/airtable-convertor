package convertor

import "testing"

func TestConvertor(t *testing.T) {
	t.Run("Run a conversion", func(t *testing.T) {
		filepath := "../tests/data.csv"
		removedChars := []string{}

		var options = make(map[string]string)
		options["eol"] = "\n"
		options["sep"] = ";"

		convertor := New(filepath, removedChars, options)
		convertor.Run()
	})
}
