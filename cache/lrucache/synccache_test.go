package lrucache

import "testing"

func Test_hashCode(t *testing.T) {
	if hashCode("12345") != 3421846044 {
		t.Error("case4 failed")
	}
	if hashCode("abcdefghijklmnopqrstuvwxyz") != 1277644989 {
		t.Error("case 5 failed")
	}
}
