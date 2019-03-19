package crypto

// BloomFilter is a probabilistic data structure for hash strings
type BloomFilter uint64

// Add adds hash string into filter
func (bf *BloomFilter) Add(hs *HBString) {
	*bf |= hs.BloomMask
}

// Test tells you if set contains item or not
// Returns false if element is definitelly not in set
// Returns true if element is probably in set
func (bf BloomFilter) Test(hs *HBString) bool {
	return hs.BloomMask == (bf & hs.BloomMask)
}

// 3 masks
var maskA = uint64(0x000000000FFFFFFF) // 7x F
var maskB = uint64(0x0000FFFFF0000000) // 5x F
var maskC = uint64(0xFFFF000000000000) // 4x F

// hashToBloom converts 64bit hash to 64bit bloom index
func hashToBloom(hash uint64) BloomFilter {

	// 23 + 22 + 19 = 64
	var a = 1 << (((hash & maskA) % 23) + 0)
	var b = 1 << (((hash & maskB) % 22) + 23)
	var c = 1 << (((hash & maskC) % 19) + 45) // 22 + 23 = 45

	return BloomFilter(a | b | c)
}
