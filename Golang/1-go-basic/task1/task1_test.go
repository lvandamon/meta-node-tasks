package main

import "testing"

func TestSingleNumber1(t *testing.T) {
	if ans := singleNumber([]int{1, 1, 2, 2, 3}); ans != 3 {
		t.Errorf("expected be 3, but %d got", ans)
	}

	if ans := singleNumber([]int{1, 1, 3, 2, 2}); ans != 3 {
		t.Errorf("expected be 3, but %d got", ans)
	}

	if ans := singleNumber([]int{3, 1, 1, 2, 2}); ans != 3 {
		t.Errorf("expected be 3, but %d got", ans)
	}
}

func TestSingleNumber2(t *testing.T) {
	t.Run("3,1,1,2,2", func(t *testing.T) {
		if singleNumber([]int{3, 1, 1, 2, 2}) != 3 {
			t.Fatal("fail")
		}

	})
}
