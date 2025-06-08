package main

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("({[]})-pos", func(t *testing.T) {
		if isValid("({[]})") != true {
			t.Fatal("fail")
		}
	})

	t.Run("({[]}-neg", func(t *testing.T) {
		if isValid("({[]}") != false {
			t.Fatal("fail")
		}
	})

}
