package main

import "testing"

func TestHelloWorld(t *testing.T) {
	// t.Fatal("not implemented")
	bf := NewBloomFilter(100)
	t.Logf("length: %d, cap: %d", len(bf.bitfield), cap(bf.bitfield))
	t.Run("check when empty", func(t *testing.T) {
		if bf.CheckValue([]byte("test")) {
			t.Errorf("bloomfilter thinks test exist while empty")
		}

	})
	t.Run("insert", func(t *testing.T) {
		 bf.CheckValue([]byte("test")
	})

	t.Run("check when has key", func(t *testing.T) {
		if !bf.CheckValue([]byte("test")) {
			t.Errorf("bloomfilter thinks test exist while empty")
		}

	})
}
