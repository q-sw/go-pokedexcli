package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestCacheExpiration(t *testing.T) {
	interval := time.Duration(time.Second * 3)
	cache := NewCache(interval)

	keyT := "Test1"
	cache.Add(keyT, []byte("valT"))

	time.Sleep(time.Second * 4)

	v, ok := cache.Get(keyT)
	fmt.Println(v)

	if ok {
		t.Errorf("Still in cache")
	}
}
