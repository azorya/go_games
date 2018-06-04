package main

import "fmt"

//  It looks like there is a rule: if you have a pointer you can get an object
//  but not every object can by used for taking it's adress
// 	it means that if our interface (myInter) is defined for a pointer like MyType below
//  an object instance can't be assigned to the interface (see _1)
//  the  "opposite case" (we have an instance of pointer to type and interface is defined for the object)
//  works fine. (see _2)

type myInter interface {
	FuncToHave(arg int) bool
}

type MyType struct {
	field1 int
}

func (mt *MyType) FuncToHave(arg int) bool {
	if mt == nil {
		fmt.Println("nil in MyType::FuncToHave")
		return false
	}

	return mt.field1 < arg
}

func Func2Call(arg myInter) {
	if arg == nil {
		fmt.Println("nil in Func2Call")
		return
	}
	arg.FuncToHave(1)
}

var _1 = func() (ret bool) {
	var I1 myInter
	var v1 *MyType // = MyType{1}
	var v2 MyType

	Func2Call(I1) // I1 == nil
	I1 = v1       // now interface != nil but the object undernith is nil
	Func2Call(I1)

	/*  !!!!!!!!!!!!!!!!!!!!
	I1 = v2 // MyType does not implement myInter !!!!
	*/
	I1 = &v2 // see the commented case above!!!!!
	Func2Call(I1)

	ret = true
	return

}()

type mytype struct {
	field int
}

func (mt mytype) FuncToHave(arg int) bool {
	return mt.field != 2
}

var _2 = func() (ret bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("my error %v", err)
			ret = false
		}
	}()
	var pmytype *mytype = &mytype{2}
	pmytype.FuncToHave(2)

	// mytype defined FuncToHave for the object not pointer
	// pmytype is a pointer and can be used (converted *) as a interface
	var I1 myInter
	I1 = pmytype

	Func2Call(I1)
	testPanic := 2
	panic(testPanic)
	return true
}()

var _2check = func() bool {
	fmt.Printf("or _2 is %v", _2)
	return true
}()

type internal int

func (a internal) FuncToHave(arg int) bool {
	a++
	return true
}

type MyTypeWithInter struct {
	// embeded field with a method satisfing our interface
	// makes our type MyTypeWithInter satisfy this interface
	internal

	notused int
}

var _3 = func() bool {
	var I1, I2 myInter //nil
	var mtwi1 MyTypeWithInter
	mtwi1 = MyTypeWithInter{10, 1}
	//    interface
	//       type    ===> myInter
	//       value
	//

	//          interface "value field" has the whole type (internal==int)

	I1 = mtwi1
	// when we copy the interface value field (internal==int) is copied
	I2 = I1
	// when we copied interfaces we realy copied underlying types too
	// so when we change the value of the variable "from which we took our original value"
	// from 10 to 11
	mtwi1.internal = 11
	// the I2 internal value representaion is not changed it is still 10
	I2.FuncToHave(2)
	//in this call (a internal) FuncToHave(arg int) bool  will get 10 not 11

	//         interface value now has a pointer to type
	mtwi1.internal = 10

	I1 = &mtwi1
	// interface "value field" now has a pointer to our type (internal == int)
	I2 = I1

	// now both I1 and I2 have "value field" pointing to the same underlaying internal==int variable
	mtwi1.internal = 11
	// when we chnage our variable we affect both I1 and I2
	I2.FuncToHave(2)
	// now call (a internal) FuncToHave(arg int) bool  will get 11 not 10 as before

	return true
}()

type internal1 int

func (a *internal1) FuncToHave(arg int) bool {
	*a++
	return true
}

type MyTypeWithInter1 struct {
	// MyTypeWithInter1 does not satisfy our interface
	// only *MyTypeWithInter1 does
	internal1

	notused int
}

var _4 = func() bool {
	var I1, I2 myInter //nil
	var mtwi1 MyTypeWithInter1
	mtwi1 = MyTypeWithInter1{10, 1}
	// now interface can't be MyTypeWithInter1
	// it can be only *MyTypeWithInter1
	I1 = &mtwi1

	I2 = I1
	// the interface value is a pointer
	// so the change below is "visible"
	mtwi1.internal1 = 11

	// but unlike the previous case we do not copy our type
	// it is a pointer to our type that implements our interface
	// func (a *internal1) FuncToHave(arg int) bool {
	// we get a pointer to   our type
	// so FuncTohave changes mtwi1.internal to 12
	I2.FuncToHave(2)
	// now mtwi1.internal == 12
	return true
}()

var _5 = func() bool {
	var I1, I2 myInter //nil

	b := (I1 == I2) //true

	var pi1, pi2 *internal

	I1, I2 = pi1, pi2

	b = I1 == I2 // true

	var pmtwi *MyTypeWithInter
	I1 = pmtwi

	b = I1 == I2 //false
	/*
	   Interface values may be compared using == and !=.
	   Two interface values are equal if both are nil,
	   or if their dynamic types are identical and their dynamic values are equal
	   according to the usual behavior of = = for that type.
	*/
	fmt.Println(b)
	return true
}()

type Interface interface {
	LessBoo() bool
}

type ProveError int32

func (pe *ProveError) LessBoo() bool { return true }

type reverse struct {
	// This embedded Interface permits Reverse to use the methods of
	// another Interface implementation.
	Interface
}

func (r reverse) LessBoo() bool { return !r.Interface.LessBoo() }

// Less returns the opposite of the embedded implementation's Less method.
//func (r reverse) Less(i, j int) bool {
//	return r.Interface.Less(j, i)
//}

// Reverse returns the reverse order for data.
func Reverse(data Interface) Interface {
	return reverse{data}
}

var _6 = func() bool {

	var pe ProveError
	var I Interface
	I = Reverse(&pe)
	I = &pe
	I.LessBoo()

	return true
}()
