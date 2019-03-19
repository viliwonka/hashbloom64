package hashbloom64

import "testing"

func TestContainsSimpleHashStringSet(T *testing.T) {

	stringSlice := []string{"sample1", "sample2", "sample3"}

	hss := NewHashStringSet(stringSlice)

	example1 := NewHashBloomString("sample1")
	example2 := NewHashBloomString("sample4")

	if !hss.Contains(example1) {
		T.Errorf("example1 should be contained!")
	}

	if hss.Contains(example2) {
		T.Errorf("example2 should't be contained!")
	}
}

func TestContainsRandomHashStringSet(T *testing.T) {

	strs := []*HBString{}

	n := 200

	// make 100 random hash strings
	for i := 0; i < n; i++ {
		strs = append(strs, NewHashBloomString(RandomStringFromRunes(32)))
	}

	hbss := NewHashStringSet(strs)

	// containment test
	for i := 0; i < n; i++ {

		if !hbss.Contains(strs[i]) {
			T.Errorf("HBStringSet should contain this element!")
		}
	}

	// non-containment test
	for i := 0; i < n; i++ {
		// string with length 31 instead of 32, should never match / be contained
		s := RandomStringFromRunes(31)

		if hbss.Contains(NewHashBloomString(s)) {
			T.Errorf("HBStringSet shouldn't contain this element!")
		}
	}

}
