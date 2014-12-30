// generated by optioner -type Example; DO NOT EDIT

// Please report issues and submit contributions at:
// http://github.com/akualab/optioner
// optioner is a project of AKUALAB INC.

package example

// Option type is used to set options in Example.
type option func(*Example) option

// Option method sets the options. Returns option to restore original values.
func (t *Example) Option(options ...option) (original option) {
	for _, opt := range options {
		original = combineOptions(original, opt(t))
	}
	return original
}

func combineOptions(a, b option) option {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	return func(t *Example) option {
		prevA := a(t)
		prevB := b(t)
		return combineOptions(prevB, prevA)
	}
}

// N sets a value for instances of type Example.
func N(o int) option {
	return func(t *Example) option {
		previous := t.N
		t.N = o
		return N(previous)
	}
}

// FSlice sets a value for instances of type Example.
func FSlice(o []float64) option {
	return func(t *Example) option {
		previous := t.FSlice
		t.FSlice = o
		return FSlice(previous)
	}
}

// Map sets a value for instances of type Example.
func Map(o map[string]int) option {
	return func(t *Example) option {
		previous := t.Map
		t.Map = o
		return Map(previous)
	}
}

// Ff sets a value for instances of type Example.
func Ff(o func(int) int) option {
	return func(t *Example) option {
		previous := t.ff
		t.ff = o
		return Ff(previous)
	}
}
