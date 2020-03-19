package test_demo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/bouk/monkey"
	jjson "github.com/json-iterator/go"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCalculate(t *testing.T) {
	monkey.Patch(json.Marshal, func(v interface{}) ([]byte, error) {
		fmt.Println("use jsoniter1")
		return jjson.Marshal(v)
	})

	monkey.Patch(json.Unmarshal, func(data []byte, v interface{}) error {
		fmt.Println("use jsoniter2")
		return jjson.Unmarshal(data, v)
	})

	Convey("start test calculate", t, func() {
		Convey("test add", func() {
			So(Add(1, 2), ShouldEqual, 3)
		})
		Convey("test sub", func() {
			So(Sub(1, 2), ShouldEqual, -1)
		})
		Convey("test mul", func() {
			So(Multiply(1, 2), ShouldEqual, 2)
		})
		Convey("test dev", func() {
			So(Devision(1, 2), ShouldEqual, 0)
		})
	})
	Convey("test more", t, func() {
		Convey("test slice equal", func() {
			// 深度比较
			So([]int{1, 2}, ShouldResemble, []int{1, 2})
		})
	})

	type Stu struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	ss := Stu{
		Name: "pengj",
		Age:  29,
	}
	stuBytes, _ := json.Marshal(ss)
	json.Unmarshal(stuBytes, &ss)

	guard := monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		return fmt.Print(a)
	})
	defer guard.Unpatch()

	fmt.Println("ss =>", ss)
	fmt.Println(string(stuBytes))
	fmt.Println("--")
	fmt.Println(string(stuBytes))

}

func TestMonkey(t *testing.T) {
	monkey.Patch(fmt.Println, func(a ...interface{}) (n int, err error) {
		s := make([]interface{}, len(a))
		for i, v := range a {
			s[i] = strings.Replace(fmt.Sprint(v), "hell", "*bleep*", -1)
		}
		return fmt.Fprintln(os.Stdout, s...)
	})
	fmt.Println("what the hell?") // what the *bleep*?
}
