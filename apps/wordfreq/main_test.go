package main

import (
	"testing"

	"solod.dev/so/strings"
)

func TestWordFrequency(t *testing.T) {
	r := strings.NewReader("hello world hello foo hello world")
	entries := countWords(&r)

	if len(entries) != 3 {
		t.Fatalf("expected 3 unique words, got %d", len(entries))
	}

	// Sorted descending: hello(3), world(2), foo(1)
	if entries[0].word != "hello" || entries[0].count != 3 {
		t.Errorf("entry 0: expected hello:3, got %s:%d", entries[0].word, entries[0].count)
	}
	if entries[1].word != "world" || entries[1].count != 2 {
		t.Errorf("entry 1: expected world:2, got %s:%d", entries[1].word, entries[1].count)
	}
	if entries[2].word != "foo" || entries[2].count != 1 {
		t.Errorf("entry 2: expected foo:1, got %s:%d", entries[2].word, entries[2].count)
	}

	freeEntries(entries)
}

func TestTopN(t *testing.T) {
	r := strings.NewReader("a a a b b c d d d d")
	entries := countWords(&r)

	if len(entries) != 4 {
		t.Fatalf("expected 4 unique words, got %d", len(entries))
	}

	// Descending by count: d(4), a(3), b(2), c(1)
	if entries[0].word != "d" || entries[0].count != 4 {
		t.Errorf("top 1: expected d:4, got %s:%d", entries[0].word, entries[0].count)
	}
	if entries[1].word != "a" || entries[1].count != 3 {
		t.Errorf("top 2: expected a:3, got %s:%d", entries[1].word, entries[1].count)
	}

	// Only top 2 should be printed in real usage
	n := 2
	if n > len(entries) {
		n = len(entries)
	}
	top := entries[:n]
	if len(top) != 2 {
		t.Errorf("expected top 2 entries, got %d", len(top))
	}

	freeEntries(entries)
}

func TestEmptyFile(t *testing.T) {
	r := strings.NewReader("")
	entries := countWords(&r)

	if len(entries) != 0 {
		t.Errorf("expected 0 entries for empty file, got %d", len(entries))
	}

	freeEntries(entries)
}

func TestSingleWord(t *testing.T) {
	r := strings.NewReader("repeat repeat repeat repeat repeat")
	entries := countWords(&r)

	if len(entries) != 1 {
		t.Fatalf("expected 1 unique word, got %d", len(entries))
	}

	if entries[0].word != "repeat" || entries[0].count != 5 {
		t.Errorf("expected repeat:5, got %s:%d", entries[0].word, entries[0].count)
	}

	freeEntries(entries)
}
