package mcache

import "sync"

const maxSize = 46

var caches [maxSize]sync.Pool

func init() {
	for i := 0; i < maxSize; i++ {
		size := 1 << i
		caches[i].New = func() any {
			return make([]byte, 0, size)
		}
	}
}

// Malloc 获取一个合适大小的[]byte
func Malloc(size int, capacity ...int) []byte {
	if len(capacity) > 1 {
		panic("too many arguments to Malloc")
	}
	var c = size
	if len(capacity) > 0 && capacity[0] > size {
		c = capacity[0]
	}
	var buf = caches[calcIndex(c)].Get().([]byte)
	buf = buf[:size]
	return buf
}

// Free 回收[]byte
func Free(buf []byte) {
	size := cap(buf)
	//大小不是2的幂, 不回收
	if !isPowerOfTwo(size) {
		return
	}
	//将buf的size置为0, 但capacity不变
	buf = buf[:0]
	//回收
	caches[bitsLen(size)-1].Put(buf)
}

// 根据size计算出合适大小的[]byte
func calcIndex(size int) int {
	if size == 0 {
		return 0
	}
	res := bitsLen(size)
	if isPowerOfTwo(size) {
		return res - 1
	}
	return res
}
