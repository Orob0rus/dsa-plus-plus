package bloom

import (
	"encoding/binary"
	"hash/fnv"
	"math"
)

type HashFunction func([]byte) uint64

type BloomFilter struct {
	store         []bool
	inputSize     int
	size          int
	hashFuncCount int
	hashFunctions []HashFunction
}

func (bf *BloomFilter) Add(val []byte) {
	for _, hashFunc := range bf.hashFunctions {
		index := hashFunc(val)
		bf.store[index] = true
	}
}

func (bf *BloomFilter) MayHaveSeen(val []byte) bool {
	for _, hashFunc := range bf.hashFunctions {
		index := hashFunc(val)
		if !bf.store[index] {
			return false
		}
	}
	return true
}

func BloomFilterCurried(inputSize int) func(size int) *BloomFilter {
	return func(size int) *BloomFilter {
		hashFuncCount := optimalHashFuncCount(size, inputSize)
		return &BloomFilter{
			store:         make([]bool, size),
			size:          size,
			inputSize:     hashFuncCount,
			hashFuncCount: hashFuncCount,
			hashFunctions: hashFunctions(size, hashFuncCount),
		}
	}
}

func optimalHashFuncCount(size, maxEntries int) int {
	sizeF := float64(size)
	maxEntriesF := float64(maxEntries)
	return int(math.Ceil(sizeF/maxEntriesF)) * int(math.Log(2))
}

func hashFunctions(size, count int) []HashFunction {
	hashFuncSlice := make([]HashFunction, count)
	for i := 0; i < count; i++ {
		hashFuncSlice[i] = func(data []byte) uint64 {
			hash := fnv.New64a()
			ibuf := make([]byte, binary.MaxVarintLen64)
			bufSize := binary.PutVarint(ibuf, int64(i))
			hash.Write(ibuf[:bufSize])
			hash.Write(data)
			return hash.Sum64() % uint64(size)
		}
	}
	return hashFuncSlice
}
