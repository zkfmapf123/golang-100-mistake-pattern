package main

import (
	"runtime"
	"testing"
)

func BenchmarkConvert1(b *testing.B) {
	// GC 통계 초기화
	runtime.GC()
	b.ReportAllocs()

	foos := make([]Foo, 1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert1(foos)
	}
}

func BenchmarkConvert2(b *testing.B) {
	// GC 통계 초기화
	runtime.GC()
	b.ReportAllocs()

	foos := make([]Foo, 1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert2(foos)
	}
}

func BenchmarkConvert3(b *testing.B) {
	// GC 통계 초기화
	runtime.GC()
	b.ReportAllocs()

	foos := make([]Foo, 1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		convert3(foos)
	}
}
