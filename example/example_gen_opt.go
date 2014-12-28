// generated by optioner -type Example; DO NOT EDIT

// Please report issues and submit contributions at:
// http://github.com/akualab/optioner
// optioner is a project of AKUALAB INC.

package example

// Option type is used to pass options to Example.
type Option func(*Example)

// Package author: use this method inside func NewExample()
// to set optional values.
func (t *Example) init(options ...Option) {
	for _, option := range options {
		option(t)
	}
}

// N sets optional value in Example.
func N(o int) Option {
	return func(t *Example) {
		t.N = o
	}
}

// FSlice sets optional value in Example.
func FSlice(o []float64) Option {
	return func(t *Example) {
		t.FSlice = o
	}
}

// Map sets optional value in Example.
func Map(o map[string]int) Option {
	return func(t *Example) {
		t.Map = o
	}
}

// Ff sets optional value in Example.
func Ff(o func(int) int) Option {
	return func(t *Example) {
		t.ff = o
	}
}
