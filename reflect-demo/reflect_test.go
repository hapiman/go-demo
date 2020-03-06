package reflect_demo

import "testing"

func TestSetValueToStruct(t *testing.T) {
	p := &Person{
		Name: "pengj   ",
		Age:  25,
	}

	startNameLen := len(p.Name)
	p = SetValueToStruct(p)
	endNameLen := len(p.Name)
	if startNameLen == endNameLen || endNameLen != 5 {
		t.Fatal("length is wrong")
	}
}
