package main

import (
	"fmt"
)

type outer struct {
	f1 int
	f2 int
}

func (o outer) Method1(arg int) {
	o.f1 += arg
}

func (o *outer) Method2(arg int) {
	o.f1 += arg
}

func mapchange(m map[string]int) {
	a := m["Hi"]
	a += 2
}

// inside of the package we just refer to the type defined in the other file
var xx = point{1, 2}

func main() {
	//var m1 map[string]int
	m1 := map[string]int{"hi": 1, "there": 2}
	if m1 == nil {
		fmt.Println("indeed")
	}
	mapchange(m1)

	// pointers vs values
	var o outer
	o.f1 = 1
	o.f2 = 2
	// do not change
	o.Method1(10)
	// does
	o.Method2(10)

	p := &o
	p.Method1(10) // syntatic sugar. it is really (*p).Method1
	return
}
