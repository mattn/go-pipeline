package pipeline

import (
	"fmt"
	"log"
	"testing"
)

func ExampleCommandPipeLine() {
	out, err := Output(
		[]string{"git", "log", "--oneline"},
		[]string{"grep", "first import"},
		[]string{"wc", "-l"},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(out))
	// Output:
	// 1
}

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
