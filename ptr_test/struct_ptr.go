package main

import (
	"fmt"
	"sort"
)

var _tested2 = proba1()

type myType struct {
	f1 int
	f2 int
	f3 string
}

// its a slice of pointers which means we do not copy /swap structs and worst
// of all strings
// myType{3} 3 means sorted by the third field
type myTypeBase []*myType

func (mt3 myTypeBase) Len() int {
	return len(mt3)
}

func (mt3 myTypeBase) Swap(i, j int) {
	mt3[i], mt3[j] = mt3[j], mt3[i]
}

func (mt3 myTypeBase) Less(i, j int) bool {
	return mt3[i].f3 < mt3[j].f3
}

func proba1() bool {
	v1 := myType{1, 2, "privet"}
	//
	//	Shortcut
	// 	If type of literal used as map key or element of array, slice or map is identical to type of key
	//  or element then such type can be omitted for brevity:
	var first = []*myType{
		{1, 2, "Hello"}, // that is why why are not writing &myType{} all the time
		{2, 3, "World"},
		{3, 1, "Alex"},
		&myType{100, 100, "hihi"},
	}

	sort.Sort(myTypeBase(first))
	ff := first[0]
	ff.f2 = 4
	fmt.Println(v1.f3, first[0].f2)
	return true
}

type typeForOtherSort struct {
	myTypeBase
}

func (tfos typeForOtherSort) Less(i, j int) bool {
	return tfos.myTypeBase[i].f2 < tfos.myTypeBase[j].f2
}

var tested2 = func() bool {
	var xxx typeForOtherSort = typeForOtherSort{
		myTypeBase{
			{1, 2, "Hello"},
			{2, 3, "World"},
			{3, 1, "Alex"},
		},
	}
	sort.Sort(xxx)
	for _, v := range xxx.myTypeBase {
		fmt.Println(v)
	}

	var first = []*myType{
		{1, 2, "Hello"},
		{2, 3, "World"},
		{3, 1, "Alex"},
	}
	sort.Sort(typeForOtherSort{first})

	return true
}()

/*
type OneMoreOtherSort myTypeBase

func (arg OneMoreOtherSort) Less(i, j int) bool {
	return arg[i].f2 < arg[j].f2
}

var tested3 = func() bool {
	var first = []*myType{
		{1, 2, "Hello"},
		{2, 3, "World"},
		{3, 1, "Alex"},
	}

	sort.Sort(OneMoreOtherSort(first))
	return true
}
*/
