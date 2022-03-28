package main

import (
	"testing"
)

func TestTrueFalse(t *testing.T) {
	got := compare(0)
	res := "Too low, try again next time!"
	if got != res {
		t.Errorf("compare(0) = %s; want %s got %s", got, res, got)
	}

	got = compare(90)
	res = "Too high, try again next time!"
	if got != res {
		t.Errorf("compare(90) = %s; want %s got %s", got, res, got)
	}

	got = compare(88)
	res = "Well Done! Your guess is correct"
	if got != res {
		t.Errorf("compare(88) = %s; want %s got %s", got, res, got)
	}

}
