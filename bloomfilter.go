package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/binary"
	"hash"
	"hash/crc32"
	"log"
)

//type hashFunction func
var hashFunctions = []hash.Hash{md5.New(), sha1.New(), crc32.NewIEEE()}

type BloomFilter struct {
	bitfield    []byte
	maxElements uint32
}

func NewBloomFilter(m uint32) *BloomFilter {
	bf := &BloomFilter{maxElements: m}

	bf.bitfield = make([]byte, m*10, m*10)

	return bf
}

func (bf BloomFilter) CheckValue(b []byte) bool {
	for _, hash := range hashFunctions {
		var b []byte
		b = hash.Sum(b)
		//_, _ = hash.Write(b)
		log.Println(len(b))
		b = b[0:4]
		index := binary.LittleEndian.Uint32(b)
		index = index % uint32(cap(bf.bitfield))
		indexByte := index / 8
		indexBit := index % 8
		targetByte := bf.bitfield[indexByte]
		if targetByte&(00000001<<indexBit) == 0 {
			return false
		}
	}
	return true
}
func (bf *BloomFilter) InsertValue(b []byte) {
	for _, hash := range hashFunctions {
		var b []byte
		hash.Sum(b)
		_, _ = hash.Write(b)
		b = b[0:4]
		index := binary.LittleEndian.Uint32(b)
		index = index % uint32(cap(bf.bitfield))
		indexByte := index / 8
		indexBit := index % 8
		bf.bitfield[indexByte] = bf.bitfield[indexByte] | (00000001 << indexBit)
	}
}
