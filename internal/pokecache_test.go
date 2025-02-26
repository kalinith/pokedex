package internal
import (
	"testing"
	"fmt"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
		{
			key: "https://example.com/otherpath",
			val: []byte("This should be interesting"),
		},
		{
			key: "https://example.com/otherpath/subpath",
			val: []byte("whatgoeshere"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestEdgeCases(t *testing.T) {
    const interval = 5 * time.Second
    cache := NewCache(interval)

    t.Run("Empty key", func(t *testing.T) {
        cache.Add("", []byte("value"))
        _, ok := cache.Get("")
        if ok {
            t.Error("expected to not find empty key")
        }
    })

    t.Run("Empty value", func(t *testing.T) {
        cache.Add("key", []byte{})
        _, ok := cache.Get("key")
        if ok {
            t.Error("expected to not find key with empty value")
        }
    })

    t.Run("Non-existent key", func(t *testing.T) {
        _, ok := cache.Get("doesnotexist")
        if ok {
            t.Error("expected false when getting non-existent key")
        }
    })

    t.Run("Overwrite existing key", func(t *testing.T) {
        key := "testkey"
        cache.Add(key, []byte("first"))
        cache.Add(key, []byte("second"))
        val, ok := cache.Get(key)
        if !ok {
            t.Error("expected to find key")
        }
        if string(val) != "second" {
            t.Error("expected to find updated value")
        }
    })
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}