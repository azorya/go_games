package main

import (
	"image/color"
	"time"
)

type point struct {
	x int
	y int
}

type point1 struct {
	x int
	y int
}

func (p *point) LongRun(arg int) {
	for i := 1; i <= arg; i++ {
		time.Sleep(5 * time.Second)
		println("LongRun")
	}
}

func (p *point) ScaleBy(arg int) {
	p.x *= arg
	p.y *= arg
}

func (p *point1) ScaleBy(arg int) {
	p.x *= arg
	p.y *= arg
}

type colorPoint struct {
	point
	color.RGBA
}

type colorPoint1 struct {
	*point1
	color.RGBA
}

var testedPtr = testPtr()

func needfuncParamInt(toCall func(int), arg int) {
	toCall(arg) //  8:8
	toCall(arg) // 16:16
}

func testPtr() bool {
	cp := colorPoint{point{1, 2}, color.RGBA{10, 20, 30, 40}}
	cp.ScaleBy(100)

	p1 := point1{y: 1, x: 1}

	cp11 := colorPoint1{&p1, color.RGBA{10, 20, 30, 40}}
	cp12 := colorPoint1{&p1, color.RGBA{11, 21, 31, 41}}
	cp12.ScaleBy(2) // 2:2
	cp11.ScaleBy(2) // 4:4

	c := color.RGBA{17, 8, 9, 0}
	c.RGBA()
	cp12.RGBA.RGBA()

	fptrLike := cp11.ScaleBy
	needfuncParamInt(fptrLike, 2)

	cp1Sb := colorPoint1.ScaleBy
	cp1Sb(cp12, 2) // 32:32
	return true
}
