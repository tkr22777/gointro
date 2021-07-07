package simplepackage

import (
	"fmt"
	"testing"
)

func TestQuote(t *testing.T) {
	want := "Hello, world."
	if got := Quote(); got != want {
		t.Errorf("Quote() = %q, want %q", got, want)
	}
	fmt.Printf("Test passed. Found:%s", want)
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
	fmt.Printf("Test passed. Found:%s", want)
}

func TestGlassy(t *testing.T) {
	want := "I can eat glass and it doesn't hurt me."
	if got := Glassy(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
	fmt.Printf("Test passed. Found:%s", want)
}
