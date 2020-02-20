// gothrift project main.go
package main

import (
	"flag"
	"fmt"
	"log"
	"reflect"
)

var (
	t int
)

func init() {
	flag.IntVar(&t, "t", 1, "please input -t=1 or -t=0")
}

type httpHeaders map[string]string

func (h httpHeaders) String() string {
	var m map[string]string = h
	return fmt.Sprintf("%s", m)
}

func main() {
	flag.Parse()

	var header httpHeaders
	header = make(httpHeaders)
	header["111"] = "1111111"
	header["222"] = "1111111"
	header["333"] = "1111111"
	header["444"] = "1111111"
	log.Println(header)

outLable:
	for i := 0; i < 5; i++ {
		fmt.Printf("外层：第%d次外层循环\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("内层：第%d次内层循环\n", j)
			continue outLable
		}
		fmt.Printf("外层：没有跳过第%d次循环\n", i)
	}

outLable1:
	for i := 0; i < 5; i++ {
		fmt.Printf("外层：第%d次外层循环\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("内层：第%d次内层循环\n", j)
			break outLable1
		}
		fmt.Printf("外层：没有跳过第%d次循环\n", i)
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("外层：第%d次外层循环\n", i)
		for j := 0; j < 5; j++ {
			fmt.Printf("内层：第%d次内层循环\n", j)
			goto outLable2
		}
		fmt.Printf("外层：没有跳过第%d次循环\n", i)
	}

outLable2:

	vv := DeferFunc2(1)
	log.Println(vv)

	// i := GetValue()
	//
	// switch i.(type) {
	// case int:
	// 	println("int")
	// case string:
	// 	println("string")
	// case interface{}:
	// 	println("interface")
	// default:
	// 	println("unknown")
	// }

	// list := new([10]int)
	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	fmt.Println(s1)

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qqq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}
	sm2 := struct {
		age int
		m   map[string]string
	}{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
	if reflect.DeepEqual(sm1, sm2) {
		fmt.Println("sm1 == sm2")
	}

	// var x *int = nil
	var yyy interface{}
	Foo(yyy)

	fmt.Println(x, y, z, k, p)

	log.Println("ok")
}

func DeferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

// func GetValue() int {
// 	return 1
// }

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
)
