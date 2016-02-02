package benchmarks_test

import (
	"testing"

	"github.com/sridharv/benchmarks"
)

func BenchmarkCallPointer(b *testing.B) {
	c := &benchmarks.Counter{}
	for i := 0; i < b.N; i++ {
		c.Add()
	}
}

func BenchmarkCallInterface(b *testing.B) {
	var a benchmarks.Adder
	a = &benchmarks.Counter{}
	for i := 0; i < b.N; i++ {
		a.Add()
	}
}

func BenchmarkCallFunctionPointer(b *testing.B) {
	c := &benchmarks.Counter{}
	fn := c.Add
	for i := 0; i < b.N; i++ {
		fn()
	}
}

func BenchmarkCallFunctionPointerInterface(b *testing.B) {
	c := &benchmarks.Counter{}
	var a benchmarks.Adder
	a = benchmarks.AdderFunc(c.Add)
	for i := 0; i < b.N; i++ {
		a.Add()
	}
}
