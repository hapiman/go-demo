package reflect_demo

import "testing"

func TestSetValueToStruct(t *testing.T) {
	p := &Person{
		Name: "pengjs   ",
		Age:  25,
		Student: Student{
			Id: "   sss  ",
		},
	}

	startNameLen := len(p.Name)
	p = TrimStructStringSpace(p).(*Person)
	endNameLen := len(p.Name)
	if startNameLen == endNameLen || endNameLen != 5 {
		t.Fatal("length is wrong")
	}
}
