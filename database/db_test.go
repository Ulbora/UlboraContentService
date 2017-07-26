package database

import "testing"

func Test_connect(t *testing.T) {
	res := Connect()
	if res != true {
		t.Fail()
	}
}
