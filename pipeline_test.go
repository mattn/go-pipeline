package pipeline

import (
	"log"
	"testing"
)

func TestCombinedOutput(t *testing.T) {
	out, err := CombinedOutput(
		[]string{"echo", "1"},
		[]string{"grep", "2"},
	)
	if err == nil {
		log.Fatal(err)
	}
	if string(out) != "" {
		log.Fatal("output is not empty.")
	}

	out, err = CombinedOutput(
		[]string{"echo", "1"},
		[]string{"rmdir", "tmptmptmp"},
	)
	if err == nil {
		log.Fatal(err)
	}
	if string(out) == "" {
		log.Fatal("output is empty.")
	}
}
