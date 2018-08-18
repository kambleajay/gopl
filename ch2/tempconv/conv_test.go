package tempconv

import "testing"

func TestCToF(t *testing.T) {
	if CToF(BoilingC) != 212 {
		t.Error("Not equal to 212")
	}
}

func TestKToC(t *testing.T) {
	if KToC(120) - Celsius(-153.15) > 0.1 {
		t.Errorf("%g != 153.15", KToC(120))
	}
}
