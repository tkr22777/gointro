package gone

import "testing"

func TestGone(t *testing.T) {
	want := "Hello, world!"
	if got := Gone(); got != want {
		t.Errorf("Gone() = %q, want %q", got, want)
	}
}

func TestQuote(t *testing.T) {
	want := "Hello, world."
	if got := Quote(); got != want {
		t.Errorf("Quote() = %q, want %q", got, want)
	}
}

func TestProverb(t *testing.T) {
	want := "Concurrency is not parallelism."
	if got := Proverb(); got != want {
		t.Errorf("Proverb() = %q, want %q", got, want)
	}
}
