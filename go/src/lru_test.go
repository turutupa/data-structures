package data_structures

import (
	"testing"
)

func TestEqualArrays(t *testing.T) {
	lru := NewLRU[string, uint32](3)

	// !!!!!!!!!!!!!! IMPORTANT !!!!!!!!!!!!!!!!
	// set this to true if you want to see the link list of the
	// LRU printed out on screen on each operation
	lru.SetVerbose(false)

	// Test initial state
	if _, exists := lru.Get("foo"); exists {
		t.Errorf("Expected Get(\"foo\") to be nil")
	}

	// Test updating key-value pairs
	lru.Update("foo", 69)
	if val, exists := lru.Get("foo"); exists && val != 69 {
		t.Errorf("Expected Get(\"foo\") to be 69")
	}

	lru.Update("bar", 420)
	if val, exists := lru.Get("bar"); exists && val != 420 {
		t.Errorf("Expected Get(\"bar\") to be 420")
	}

	lru.Update("baz", 1337)
	if val, exists := lru.Get("baz"); exists && val != 1337 {
		t.Errorf("Expected Get(\"baz\") to be 1337")
	}

	lru.Update("ball", 69420)
	if val, exists := lru.Get("ball"); exists && val != 69420 {
		t.Errorf("Expected Get(\"ball\") to be 69420")
	}

	// Test LRU eviction
	if _, exists := lru.Get("foo"); exists {
		t.Errorf("Expected Get(\"foo\") to not exist")
	}

	if val, exists := lru.Get("bar"); exists && val != 420 {
		t.Errorf("Expected Get(\"bar\") to be 420")
	}

	lru.Update("foo", 69)
	if val, exists := lru.Get("bar"); exists && val != 420 {
		t.Errorf("Expected Get(\"bar\") to be 420")
	}

	if val, exists := lru.Get("foo"); exists && val != 69 {
		t.Errorf("Expected Get(\"foo\") to be 69")
	}

	// Test eviction due to LRU access pattern
	if _, exists := lru.Get("baz"); exists {
		t.Errorf("Expected Get(\"baz\") to not exist")
	}
}
