package main

import "fmt"

// some commments about Pic
func Pic(dx, dy int) [][]uint8 {
	fmt.Println(dx, dy)
	yy := make([][]uint8, dy)
	for i, _ := range yy {
		xx := make([]uint8, dx)
		for j, _ := range xx {
			xx[j] = uint8(10)
		}
		yy[i] = append(yy[i], xx...)
		//fmt.Println(line)
	}
	return yy
}

/*
type alex struct {
	f1 int8
	f2 int8
}

func (al *alex) method1(arg int8) {
	al.f1 = arg
}
*/

type Fake struct {
	n1 *int
}

func (f Fake) Write(p []byte) (n int, err error) {
	*f.n1++
	n = *f.n1
	err = nil
	return
}

func main() {
	fmt.Println(Pic(2, 3))
	//pic.Show(Pic)
	//	var aaa alex
	//pa := &aaa
	//	aaa.method1(int8(32))
	var a int = 11
	f := Fake{&a}
	n, myError := fmt.Fprint(f, "hello")
	fmt.Println(n, myError)
	n, myError = fmt.Fprint(f, "hello")
}
