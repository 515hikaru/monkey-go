package object

import "testing"

func TestStringHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello, World"}
	hello2 := &String{Value: "Hello, World"}
	diff1 := &String{Value: "My name is Johnny"}
	diff2 := &String{Value: "My name is Johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with same content have different hash key.")
	}
	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with same content have different hash key.")
	}
	if hello1.HashKey() == diff2.HashKey() {
		t.Errorf("strings with different content have same hash key.")
	}

}

func TestIntegerHashKey(t *testing.T) {
	one1 := &Integer{Value: 1}
	one2 := &Integer{Value: 1}
	two1 := &Integer{Value: 2}
	two2 := &Integer{Value: 2}
	if one1.HashKey() != one2.HashKey() {
		t.Errorf("integers with same content have different hash key.")
	}
	if two1.HashKey() != two2.HashKey() {
		t.Errorf("integers with same content have different hash key.")
	}
	if one1.HashKey() == two1.HashKey() {
		t.Errorf("integers with different content have same hash key.")
	}
}

func TestBooleanHashKey(t *testing.T) {
	true1 := &Boolean{Value: true}
	true2 := &Boolean{Value: true}
	false1 := &Boolean{Value: false}
	false2 := &Boolean{Value: false}

	if true1.HashKey() != true2.HashKey() {
		t.Errorf("booleans with same content have different hash key.")
	}

	if false1.HashKey() != false2.HashKey() {
		t.Errorf("booleans with same content have different hash key.")
	}
	if true1.HashKey() == false1.HashKey() {
		t.Errorf("booleans with different content have same hash key.")
	}
}
