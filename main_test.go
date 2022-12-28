package main

import "testing"

func TestDivision(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("测试没通过")
	} else {
		t.Log("测试通过")
	}
}
