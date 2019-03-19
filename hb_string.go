package hashbloom64

import (
	"fmt"

	"github.com/cespare/xxhash"
)

// HBString (or longer HashBloomString) is performant structure for comparing strings
// used for fast comparing, slow on creation though
type HBString struct {
	Strg      string
	Hash      uint64
	BloomMask BloomFilter
}

// Equals check for equality of strings
func (hs *HBString) Equals(otherHs *HBString) bool {
	// compare hash, if hash isn't equal then string definitely isn't equal
	if hs.Hash == otherHs.Hash {
		// do actual comparsion
		return hs.Strg == otherHs.Strg
	}
	return false
}

// String generates content of HashString object
func (hs *HBString) String() string {
	return fmt.Sprintf("\n| HashString\n| Strg: %v\n| Hash:%v\n| Mask:%v\n", hs.Strg, fmt.Sprintf("\t%0.64bb", hs.Hash), fmt.Sprintf("\t%0.64bb\n", hs.BloomMask))
}

// NewHashBloomString creates new HashString from existing string object, uses xxhash
func NewHashBloomString(str string) *HBString {

	hs := HBString{Strg: str}
	h := xxhash.New()
	h.Write([]byte(str))

	hs.Hash = h.Sum64()
	// Calculate bloom mask
	hs.BloomMask = hashToBloom(hs.Hash)

	return &hs
}
