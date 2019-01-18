package days

import (
	"fmt"
	"testing"
)

func TestDelta(t *testing.T) {
	cases := []struct {
		y1, m1, d1 int
		y2, m2, d2 int
		want       int
	}{
		{2000, 1, 1, 2000, 1, 1, 0},
		{2000, 1, 1, 2000, 1, 2, 1},
		{2000, 1, 2, 2000, 1, 1, -1},
		{2000, 1, 1, 2000, 2, 1, 31},
		{2000, 1, 1, 2000, 3, 1, 60},
		{1999, 12, 31, 2000, 1, 1, 1},
		{2003, 1, 1, 2004, 1, 1, 365},
		{2003, 1, 1, 2005, 1, 1, 731},
		{1900, 2, 1, 1900, 3, 1, 28},
		{1904, 2, 1, 1904, 3, 1, 29},
		{2000, 2, 1, 2000, 3, 1, 29},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
			d := Delta(c.y1, c.m1, c.d1, c.y2, c.m2, c.d2)
			if d != c.want {
				t.Errorf("got %d, want %d", d, c.want)
			}
		})
	}
}

func TestCountLeapYears(t *testing.T) {
	cases := []struct {
		y1, y2 int
		want   int
	}{
		{2002, 2002, 0},
		{2002, 2003, 0},
		{2002, 2004, 1},
		{2002, 2005, 1},
		{2001, 2005, 1},
		{2000, 2005, 2},
		{1999, 2005, 2},
		{1900, 1901, 0},
		{1900, 1904, 1},
		{1895, 1905, 2},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
			l := CountLeapYears(c.y1, c.y2)
			if l != c.want {
				t.Errorf("got %d, want %d", l, c.want)
			}
		})
	}
}

func TestDeltaYD(t *testing.T) {
	cases := []struct {
		y1, m1, d1   int
		y2, m2, d2   int
		wantY, wantD int
	}{
		{2000, 1, 1, 2000, 1, 1, 0, 0},
		{2000, 1, 1, 2000, 1, 2, 0, 1},
		{2000, 1, 2, 2000, 1, 1, 0, -1},
		{2000, 1, 1, 2000, 2, 1, 0, 31},
		{2000, 1, 1, 2000, 3, 1, 0, 60},
		{1999, 12, 31, 2000, 1, 1, 0, 1},
		{2003, 1, 1, 2004, 1, 1, 1, 0},
		{2003, 1, 1, 2005, 1, 1, 2, 0},
		{1900, 2, 1, 1900, 3, 1, 0, 28},
		{1904, 2, 1, 1904, 3, 1, 0, 29},
		{2000, 2, 1, 2000, 3, 1, 0, 29},
	}
	for i, c := range cases {
		t.Run(fmt.Sprintf("case_%d", i+1), func(t *testing.T) {
			dy, dd := DeltaYD(c.y1, c.m1, c.d1, c.y2, c.m2, c.d2)
			if dy != c.wantY || dd != c.wantD {
				t.Errorf("got %d, %d; want %d, %d", dy, dd, c.wantY, c.wantD)
			}
		})
	}
}
