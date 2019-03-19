package hashbloom64

import (
	"math/rand"
	"strconv"
	"testing"
)

var alphabetRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func RandomStringFromRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphabetRunes[rand.Intn(len(alphabetRunes))]
	}
	return string(b)
}

func TestEqualsSimpleHashString(T *testing.T) {

	example1 := NewHashBloomString("sample")
	example2 := NewHashBloomString("sample2")
	example3 := NewHashBloomString("sample")

	if example1.Equals(example2) {
		T.Errorf("example1 shouldn't be equal to example2")
	}

	if !example1.Equals(example3) {
		T.Errorf("example1 should be equal to example3")
	}
}

func TestEqualsRandomHashString(T *testing.T) {

	strs := []*HBString{}

	n := 2000

	// make 200 random hash strings
	for i := 0; i < n; i++ {
		strs = append(strs, NewHashBloomString(RandomStringFromRunes(69)))
	}

	// cross-product comparison
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

			exact := (strs[i].Strg == strs[j].Strg)

			if exact != strs[i].Equals(strs[j]) {
				T.Errorf("HashString %v mismatched with %v comparison", strs[i].Strg, strs[j].Strg)
			}
		}
	}
}

// TestStringerInterfaceHashString tests stringer interface (String())
func TestStringerInterfaceHashString(T *testing.T) {

	for i := 0; i < 24; i++ {
		hs := NewHashBloomString("sample_string" + string(strconv.Itoa(i)))
		T.Logf("%v", hs)
	}

	T.Error("no error - just here to write out to logs")
}
