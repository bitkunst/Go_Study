package main

import "fmt"

func main() {
	var a int = 3 // 64bit 컴퓨터일 경우 8byte 정수타입
	var b int // 초기값 생략시 기본값 0
	var c = 4 
	d := 5

	var e = "Hello"
	f := 3.14 // 실수값의 기본 타입은 float64

	fmt.Println(a, b, c, d, e, f)

	// 타입 변환
	g := 2 // int
	var h float64 = 2.3

	var i int = int(h) // 2.3 -> 2
	j := float64(g) * h

	var k int64 = 7
	l := g * int(k)
	fmt.Println(g, h, i, j, k, l)
}