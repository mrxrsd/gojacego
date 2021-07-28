package cache

import "testing"

func TestCache(t *testing.T) {
	c := NewCache()

	if _, found := c.Get("key"); found {
		t.Errorf("cache should be nil for 'key'")
	}

	c.Add("key", "1")

	if item, found := c.Get("key"); found {
		if item.(string) != "1" {
			t.Errorf("expected: 1, got: %s", item.(string))
		}
	} else {
		t.Errorf("cache should not be nil")
	}

	if item, found := c.Get("key"); found {
		if item.(string) != "1" {
			t.Errorf("expected: 1, got: %s", item.(string))
		}
	} else {
		t.Errorf("cache should not be nil")
	}

	if _, found := c.Get("key2"); found {
		t.Errorf("cache should be nil for 'key2'")
	}

	ret := c.Add("key", "2")

	if ret == nil {
		t.Errorf("cannot overwrite cache entry")
	}
}
