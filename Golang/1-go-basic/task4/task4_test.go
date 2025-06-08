package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {

	t.Run("len =0", func(t *testing.T) {
		strs := []string{}
		if longestCommonPrefix(strs) != "" {
			t.Fatal("fail")
		}
	})

	t.Run("len = 1", func(t *testing.T) {
		strs := []string{"flower"}
		if longestCommonPrefix(strs) != "flower" {
			t.Fatal("fail")
		}
	})

	t.Run("length >=2,result=fl", func(t *testing.T) {
		strs := []string{"flower", "flow", "flight"}
		if longestCommonPrefix(strs) != "fl" {
			t.Fatal("fail")
		}
	})

	t.Run("length >=2,result=dog", func(t *testing.T) {
		strs := []string{"dog1", "dog2", "dog3", "dog4"}
		if longestCommonPrefix(strs) != "dog" {
			t.Fatal("fail")
		}
	})

	t.Run("length >=2,result=d", func(t *testing.T) {
		strs := []string{"da1", "db2", "dc3", "dd4"}
		if longestCommonPrefix(strs) != "d" {
			t.Fatal("fail")
		}
	})

}
