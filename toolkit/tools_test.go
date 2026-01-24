package toolkit

import "testing"

func TestTool_RandomString(t *testing.T) {
	var testTools Tools
	s := testTools.RandomString(10)
	if len(s) != 10 {
		t.Errorf("RandomString() expected length of 10, got %d", len(s))
	}
}
