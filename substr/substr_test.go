package main

import (
	"fmt"
	"testing"
)

func TestFindNoDiffStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		{"abcabcabca", 3},
		{"pkgqweqwe", 6},
		{"", 0},
		{"b", 1},
		{"bbbbbb", 1},
		{"abcdefghijklmn", 14},
	}
	for _, tt := range tests {
		actual := findNoDiffStr(tt.s)
		if actual != tt.ans {
			fmt.Printf("%s: expect:%d,actual:%d", tt.s, tt.ans, actual)
		}
	}
}

func BenchmarkFindNoDiffStr(b *testing.B) {
	s := "abcdefghijklmn"
	ans := 14
	for i := 0; i <= b.N; i++ {
		actual := findNoDiffStr(s)
		if actual != ans {
			fmt.Printf("%s: expect:%d,actual:%d", s, ans, actual)
		}

	}

}
