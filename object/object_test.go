package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{"Hello World"}
	hello2 := &String{"Hello World"}
	diff1 := &String{"My name is johnyy"}
	diff2 := &String{"My name is johnyy"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("Strings with same content have different hash keys")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("Strings with different content have same hash keys")
	}
}
