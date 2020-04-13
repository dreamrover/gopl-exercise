// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

const N = 32 << ((^uint(0)) >> 63)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/N, uint(x%N)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/N, uint(x%N)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Len() (len int) {
	var i uint
	for _, word := range s.words {
		for i = 0; i < N; i++ {
			if word&(1<<i) != 0 {
				len++
			}
		}
	}
	return
}

func (s *IntSet) Remove(x int) {
	word, bit := x/N, uint(x%N)
	if word >= len(s.words) {
		return
	}
	s.words[word] &= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	var t IntSet
	copy(t.words, s.words)
	return &t
}

func (s *IntSet) AddAll(ints ...int) {
	if len(ints) == 0 {
		return
	}
	for _, x := range ints {
		s.Add(x)
	}
}

func (s *IntSet) InsectWith(t *IntSet) *IntSet {
	var u IntSet
	if len(s.words) == 0 || len(t.words) == 0 {
		return &u
	}
	var size int
	var a, b = s.words, t.words
	if len(s.words) < len(t.words) {
		size = len(s.words)
	} else {
		size = len(t.words)
		a, b = b, a
	}
	copy(u.words, b[:size])
	for i, word := range a {
		u.words[i] &= word
	}
	return &u
}

func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	var u IntSet
	if len(s.words) == 0 {
		return &u
	}
	copy(u.words, s.words)
	w := t.words
	if len(s.words) < len(t.words) {
		w = w[:len(s.words)]
	}
	for i, word := range w {
		u.words[i] &^= word
	}
	return &u
}

func (s *IntSet) SymmeticDifference(t *IntSet) *IntSet {
	var u IntSet
	var a, b = s.words, t.words
	if len(a) == 0 && len(b) == 0 {
		return &u
	}
	if len(a) < len(b) {
		a, b = b, a
	}
	copy(u.words, a)
	for i, word := range b {
		u.words[i] ^= word
	}
	return &u
}

func (s *IntSet) Elems() []int {
	if len(s.words) == 0 {
		return nil
	}
	e := make([]int, s.Len(), s.Len())
	var j uint
	for i, w := range s.words {
		for j = 0; j < N; j++ {
			if w&(1<<j) != 0 {
				e = append(e, i*N+int(j))
			}
		}
	}
	return e
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < N; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", N*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string
