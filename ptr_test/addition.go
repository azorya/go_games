package main

import (
	"fmt"
)

// ZInt some fake comment
type ZInt interface {
	Boo() int
}

// T1 fake comment
type T1 int

// T2 fake comment
type T2 int

// Boo fake comment
func (t1 T1) Boo() int {
	return int(t1)
}

// Boo fake comment
func (t2 T2) Boo() int {
	return int(t2)
}

func (t1 T1) String() string {
	return fmt.Sprintf("%d", int(t1))
}

func (t2 T2) String() string {
	return fmt.Sprintf("%d", int(t2))
}

func test1(arg bool) ZInt {
	if arg {
		var r T1 = 2
		return r
	}

	var r T2 = 3
	return r

}

var tested = func() bool {
	var t1 T1 = 2
	var t2 T2 = 3
	xx.x = 4
	t1 = T1(t2)
	ind1 := T1.String
	ind1(t1)
	zi := test1(true)
	zi.Boo()
	fmt.Println(t1.String(), t2.String())
	return true
}()
