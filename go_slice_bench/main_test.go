package main

import "testing"

func BenchmarkMethod1(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		Method1()
	}
}

func BenchmarkMethod2(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		Method2()
	}
}

func BenchmarkMethod3(b *testing.B) {
	b.ResetTimer()
	for range b.N {
		Method3()
	}
}
