//
// A type assertion is an operation applied to an interface value.
// Syntactically, it looks like x.( T),
// where x is an expression of an interface type and T is a type, called the “asserted” type.
// A type assertion checks that the dynamic type of its operand matches the asserted type.

package main

import (
	"errors"
	"fmt"
	"os"
	"syscall"
)

// First, if the asserted type T is a concrete type,
// then the type assertion checks whether x’s dynamic type is identical to T.

type myType1 int
type myType2 int

// here we satisfy myInter (defined in interface_ptr)
// just to show that iside of the package we see everything
func (mt1 myType1) FuncToHave(arg int) bool {
	var tmp = int(mt1)
	return tmp == arg
}

func (mt1 myType2) FuncToHave(arg int) bool {
	var tmp = int(mt1)
	return tmp == arg
}

var __1 = func() (ret bool) {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Printf("%T", r)
			ret = false
		}
	}()

	pmt1 := new(myType1)
	*pmt1 = 2

	// our interface has a pointer as it's value
	var I1 myInter = pmt1
	mt1 := I1.(*myType1)

	// now it has a value as a value
	I1 = *pmt1
	mt11 := I1.(myType1)
	/*
		fmt.Printf("%T", mt1) // main.myType1
	*/
	mt2 := myType2(2)
	I1 = mt2
	// if checking the (wrong) type like this
	if mt22, ok := I1.(myType1); ok {
		mt22 = 3
		_ = mt22
	}
	// and once again to use panic/recover
	// note that mt22 that beloned to if is not visible here
	mt22 := I1.(myType1)
	_ = mt22
	_ = mt11

	_ = mt1
	return true
}()

// Second, if instead the asserted type T is an interface type,
// then the type assertion checks whether x’s dynamic type satisfies T.

type myInter2 interface {
	Func2Call(int) bool
}

// myType1 is implementing also myInter2
// myType2 is not
func (mt1 myType1) Func2Call(arg int) bool {
	return int(mt1) == arg
}

var __2 = func() bool {
	var I1 myInter = myType1(3)
	var I2 myInter2
	I2 = I1.(myInter2)
	I2.Func2Call(2)

	I1 = myType2(3)
	if I2, ok := I1.(myInter2); ok { // ok is false
		I2.Func2Call(3)
	}
	return true
}()
var __3 = func() bool {
	r := os.IsNotExist(e)
	// to be understood!!!!
	// how can one assign a Errno uintptr to interface?
	aa = 2
	r = os.IsNotExist(aa)
	return r
}()

var aa syscall.Errno

// look at error.go
var e = errors.New("file does not exist")
