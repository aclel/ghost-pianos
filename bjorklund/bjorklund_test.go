package bjorklund

import "testing"

func TestBjorklund(t *testing.T) {
	cases := []struct {
		stepsIn int
		pulsesIn int
		want []int
	}{
		{5, 2, []int{1, 0, 1, 0, 0}},
		{4, 3, []int{1, 1, 1, 0}},
		{5, 3, []int{1, 0, 1, 0, 1}},
		{7, 3, []int{1, 0, 1, 0, 1, 0, 0}},
		{8, 3, []int{1, 0, 0, 1, 0, 0, 1, 0}},
		{13, 5, []int{1, 0, 0, 1, 0, 1, 0, 0, 1, 0, 1, 0, 0}},
		{7, 4, []int{1, 0, 1, 0, 1, 0, 1}},
		{9, 4, []int{1, 0, 1, 0, 1, 0, 1, 0, 0}},
		{11, 4, []int{1, 0, 0, 1, 0 ,0, 1, 0, 0, 1, 0}},
		{6, 5, []int{1, 1, 1, 1, 1, 0}},
		//{7, 5, []int{1, 0, 1, 1, 0, 1, 1}},
		{8, 5, []int{1, 0, 1, 1, 0, 1, 1, 0}},
		{9, 5, []int{1, 0, 1, 0, 1, 0, 1, 0, 1}},
		{11, 5, []int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 0}},
		{12, 5, []int{1, 0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 0}},
		{16, 5, []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0, 0}},
	}

	for _, c := range cases {
		got := Bjorklund(c.stepsIn, c.pulsesIn)
		if !sliceEqual(got, c.want) {
			t.Errorf("Bjorklund(%d, %d) == %v, want %v", c.stepsIn, c.pulsesIn, got, c.want)
		}
	}
}

func sliceEqual(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}