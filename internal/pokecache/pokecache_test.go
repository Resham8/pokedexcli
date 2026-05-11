package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}

}

func TestAddGetCache(t *testing.T) {
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "key1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "key2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}

	for _, cas := range cases {
		cache := NewCache(time.Millisecond)
		cache.Add(cas.inputKey, cas.inputVal)

		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Errorf("%v not found", cas.inputKey)
			continue
		}

		if string(actual) != string(cas.inputVal) {
			t.Errorf("%v doesn't match %s", string(cas.inputVal), string(actual))
			continue
		}
	}

}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)
	inputKey := "key1"
	inputVal := []byte("val1")
	cache.Add(inputKey, inputVal)
	time.Sleep(interval + time.Microsecond)

	_, ok := cache.Get(inputKey)
	if ok {
		t.Errorf("%s should have been reaped",inputKey)
	} 
}
