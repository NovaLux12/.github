package main

import "testing"

func TestOk(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("math broke")
	}
}
