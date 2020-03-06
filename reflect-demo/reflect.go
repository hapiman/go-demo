package reflect_demo

import (
	"fmt"
	"reflect"
	"strings"
)

type Student struct {
	Id string `json:"id"`
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Student
}

func TrimStructStringSpace(p interface{}) interface{} {
	vv := reflect.ValueOf(p).Elem()
	tt := reflect.TypeOf(p).Elem()
	for i := 0; i < vv.NumField(); i++ {
		f := vv.Field(i)
		switch f.Kind() {
		case reflect.String:
			old := f.String()
			new := strings.TrimSpace(old)
			key := tt.Field(i).Name
			fmt.Printf("old: %s, new: %s, key: %s", old, new, key)
			vv.FieldByName(key).Set(reflect.ValueOf(new))
		}
	}
	return p
}
