package challenges

import (
	"testing"
)

var lru = LRU[string, int]{
	capacity:      3,
	lookup:        make(map[string]*Node[int]),
	reverseLookup: make(map[*Node[int]]string),
}

func TestLRU(t *testing.T) {
	if _, ok := lru.get("foo"); ok {
		t.Error("Key shouldn't exist")
	}
	lru.update("foo", 69)
	if value, ok := lru.get("foo"); ok {
		if value != 69 {
			t.Errorf("Expected %v but got %v", 69, value)
		}
	}

	lru.update("bar", 122)
	if value, ok := lru.get("bar"); ok {
		if value != 122 {
			t.Errorf("Expected %v but got %v", 122, value)
		}
	}

	lru.update("baz", 10)
	if value, ok := lru.get("baz"); ok {
		if value != 10 {
			t.Errorf("Expected %v but got %v", 10, value)
		}
	}

	lru.update("ball", 1000)
	if value, ok := lru.get("ball"); ok {
		if value != 1000 {
			t.Errorf("Expected %v but got %v", 1000, value)
		}
	}

	if _, ok := lru.get("foo"); ok {
		t.Error("Key shouldn't exist")
	}

	if value, ok := lru.get("bar"); ok {
		if value != 122 {
			t.Errorf("Expected %v but got %v", 122, value)
		}
	}

	lru.update("foo", 69)

	if value, ok := lru.get("bar"); ok {
		if value != 122 {
			t.Errorf("Expected %v but got %v", 122, value)
		}
	}

	if value, ok := lru.get("foo"); ok {
		if value != 69 {
			t.Errorf("Expected %v but got %v", 69, value)
		}
	}

	if _, ok := lru.get("baz"); ok {
		t.Error("Key shouldn't exist")
	}
}
