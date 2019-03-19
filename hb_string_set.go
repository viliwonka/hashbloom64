package crypto

// HBStringSet is slice of hashstrings
type HBStringSet struct {
	Slice  []*HBString
	Filter BloomFilter
}

// AddHashString adds new hash string in set, if not already contained
func (hbss *HBStringSet) AddHashString(hs *HBString) bool {

	if hbss.Contains(hs) {
		return false
	}
	// add hash strings
	hbss.Slice = append(hbss.Slice, hs)
	hbss.Filter.Add(hs)
	return true
}

// AddString adds new string in set, if not already contained
func (hbss *HBStringSet) AddString(s string) bool {

	hs := NewHashBloomString(s)

	if hbss.Contains(hs) {
		return false
	}

	// add hash strings
	hbss.Slice = append(hbss.Slice, hs)
	hbss.Filter.Add(hs)

	return true
}

// Contains checks if HashStringSlice contains hash string
func (hbss *HBStringSet) Contains(hs *HBString) bool {

	if hbss.Filter.Test(hs) {
		for _, v := range hbss.Slice {
			if v.Equals(hs) {
				return true
			}
		}
	}

	return false
}

// NewHashStringSet makes new Hash set from []string, []HashString or []*HashString
func NewHashStringSet(slice interface{}) *HBStringSet {

	hss := HBStringSet{}

	switch slice.(type) {

	case []string:

		stringSlice := slice.([]string)
		for _, s := range stringSlice {
			hss.AddString(s)
		}
		break

	case []HBString:

		hsSlice := slice.([]HBString)
		for _, hs := range hsSlice {
			hss.AddHashString(&hs)
		}
		break

	case []*HBString:

		hss.Slice = slice.([]*HBString)
	}

	// calculate bloom mask
	var maskUnion BloomFilter
	for _, hs := range hss.Slice {
		maskUnion |= hs.BloomMask
	}

	hss.Filter = maskUnion

	return &hss
}
