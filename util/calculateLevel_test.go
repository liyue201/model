package util

import (
	"testing"
)

func TestCalculateLevel(t *testing.T) {
	level, point1, point2 := CalculateUserLevel(600)
	println(level)
	println(point1)
	println(point2)

	if level != 3 {
		t.Fatalf(`error`)
	}
	if point1 != 399 {
		t.Fatalf(`error`)
	}
	if point2 != 619 {
		t.Fatalf(`error`)
	}
}
