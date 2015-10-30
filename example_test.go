package pipeline

import (
	"fmt"
	"log"
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
