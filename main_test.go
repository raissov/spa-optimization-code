package main

import (
	"testing"
)



func BenchmarkSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}

func BenchmarkFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solution()
	}
}