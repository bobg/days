package days

import "testing"

func TestDelta(t *testing.T) {
	cases := []struct {
		y1, m1, d1, y2, m2, d2 int
		want                   int
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
	for _, c := range cases {
		d := Delta(c.y1, c.m1, c.d1, c.y2, c.m2, c.d2)
		if d != c.want {
			t.Errorf("case %d/%d/%d-%d/%d/%d: got %d, want %d", c.y1, c.m1, c.d1, c.y2, c.m2, c.d2, d, c.want)
		}
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
	for _, c := range cases {
		l := CountLeapYears(c.y1, c.y2)
		if l != c.want {
			t.Errorf("case %d-%d: got %d, want %d", c.y1, c.y2, l, c.want)
		}
	}
}
