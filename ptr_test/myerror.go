// Here we use error interface as an example to show how interfaces are compared
//
//

package main

import (
	"strconv"
)

type myStrange struct {
	int
	bool
}

func (ms myStrange) Error() string {
	i, _ := ms.int, ms.bool
	return strconv.Itoa(i)
}

var msg = myStrange{10, true}

func getSame(same bool) error {
	if same {
		return &msg
	}

	return &myStrange{10, true}
}

func t1() bool {
	s1 := myStrange{10, true}
	s2 := myStrange{10, true}

	// err1 below has a pointer as an interface value
	var err1 error = &s1
	// err2 has a myStrange object as an interface value
	var err2 error = s1
	b := err1 == err2 // false

	// err2 has a pointer (as err1) but the value of this pointers are different
	err2 = &s2
	b = err1 == err2 // false

	// both err1 & err2 has values and these values are the same
	err1 = s1
	err2 = s2
	b = err1 == err2 // true

	// both have pointers and pointers are the same
	err1 = getSame(true)
	err2 = getSame(true)
	b = err2 == err2 // true

	// both have pointers and pointers are different (yes it is not C!)
	err1 = getSame(false)
	err2 = getSame(false)
	b = err1 == err2 // false

	_ = b

	return true
}

var bt = t1()
