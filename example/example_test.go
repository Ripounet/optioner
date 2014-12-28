package example

import "testing"

func TestExample(t *testing.T) {

	myFunc := func(n int) int { return 2 * n }
	ex := NewExample("test", N(22), Ff(myFunc))

	if ex.N != 22 {
		t.Errorf("N is %d, expected 22", ex.N)
	}

	if ex.ff(10) != 20 {
		t.Errorf("ff(10) is %d, expected 20", ex.ff(10))
	}

	if cap(ex.FSlice) != 100 {
		t.Errorf("FSlice cap is %d, expected 100", cap(ex.FSlice))
	}
}