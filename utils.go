package mcache

import "math/bits"

// x有多少个二进制位
func bitsLen(x int) int {
	return bits.Len(uint(x))
}

// 是否是2的幂
func isPowerOfTwo(x int) bool {
	return (x & (-x)) == x
}
