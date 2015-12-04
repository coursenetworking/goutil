package goutil

import "testing"

func TestT(t *testing.T) {
	if s := SubStr("abcdef", 0, 1); s != "a" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", 1, 2); s != "bc" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", -2, 1); s != "e" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", -2, 0); s != "ef" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", 1, -2); s != "bcd" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", -3, -2); s != "d" {
		t.Errorf("not match [%s] ", s)
	}

	// please note
	if s := SubStr("abcdef", -2, -4); s != "" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", -20, -4); s != "ab" {
		t.Errorf("not match [%s] ", s)
	}

	if s := SubStr("abcdef", -20, -10); s != "" {
		t.Errorf("not match [%s] ", s)
	}
}
