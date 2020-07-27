package pipeline

import (
	"bytes"
	"context"
	"log"
	"testing"
)

func TestOutput(t *testing.T) {
	out, err := Output(
		[]string{"echo", "hello"},
		[]string{"wc", "-c"},
	)
	if err != nil {
		t.Fatal(err)
	}

	expect := []byte("6\n")
	if bytes.Compare(out, expect) != 0 {
		t.Errorf("expect %s but %s", expect, out)
	}
}

func TestOutputContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := OutputContext(
		ctx,
		[]string{"echo", "hello"},
		[]string{"wc", "-c"},
	)
	if err != context.Canceled {
		t.Fatal("err should be context.Canceled")
	}
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
