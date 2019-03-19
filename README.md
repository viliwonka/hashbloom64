# Hash Bloom 64

This repository contains two structures:

* HBString,
* HBStringSet

HBString is short for HashBloomString, and as you have guessed, HBStringSet is HashBloomStringSet. 

HBString contains string, hash and bloom filter. Hash is used for fast comparison between strings, and bloom is used for fast testing if string is inside set or not (set = slice in golang). HBString structure is also space-efficient, since it uses only 64 bits for hash and 64 bits for bloom.


H = Hash

String bytes are hashed by xxHash. I did short research and xxHash was best fit for this project, it's very fast, very random, and produces 64 bits.

B = Bloom

A quote from Wikipedia, that already describes what Bloom filter is:

```A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set".```

The Bloom filter is created from Hash, and is also again 64 bits long.

Usage: 

For fast comparision between strings, use `HBString` and it's method `Equals`, which returns true/false:

```go

    example1 := NewHashBloomString("sample")
    example2 := NewHashBloomString("sample2")
    
    result := example1.Equals(other example2)

``` 

For fast checking for containment in sets, use `HBStringSet`

```go

    stringSlice := []string{"sample1", "sample2", "sample3"}

	hss := NewHashStringSet(stringSlice)

	example1 := NewHashBloomString("sample1")
	example2 := NewHashBloomString("sample4")

    // return true
	result := hss.Contains(example1)
    // returns false
    result = hss.Contains(example2)

```