package convertor

import (
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("Dispatch mapped channels", func(t *testing.T) {
		channels := make(map[string]chan []string)

		channels["read"] = make(chan []string)
		channels["write"] = make(chan []string)

		dispatchChannels(channels)

	})

	t.Run("Create a new parser", func(t *testing.T) {
		addChars := []string{}

		channels := make(map[string]chan []string)

		channels["read"] = make(chan []string)
		channels["write"] = make(chan []string)

		parser := NewParser(addChars, channels)

		if got, want := parser.ToBeRemoved, unWantedChars; len(got) != len(want) {
			t.Errorf("GOT: %v; WANT: %v", got, want)
		}
	})

}
