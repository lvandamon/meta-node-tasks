package main

import "testing"

func TestIsPalindrome(t *testing.T) {

	t.Run("1-true", func(t *testing.T) {
		if isPalindrome(1) != true {
			t.Fatal("fail")
		}
	})

	t.Run("121-true", func(t *testing.T) {
		if isPalindrome(121) != true {
			t.Fatal("fail")
		}
	})

	t.Run("1221-true", func(t *testing.T) {
		if isPalindrome(1221) != true {
			t.Fatal("fail")
		}
	})

	t.Run("-121-false", func(t *testing.T) {
		if isPalindrome(-121) != false {
			t.Fatal("fail")
		}
	})
	t.Run("10-false", func(t *testing.T) {
		if isPalindrome(10) != false {
			t.Fatal("fail")
		}
	})
}
