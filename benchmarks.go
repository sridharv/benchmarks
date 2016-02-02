package benchmarks

type Adder interface {
	Add()
}

type Counter struct {
	Count int
}

func (c *Counter) Add() {
	c.Count++
}

type AdderFunc func()

func (a AdderFunc) Add() { a() }
