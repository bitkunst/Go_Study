package main

import "fmt"

func main() {
	var a int = 10 // integer(정수) 타입 변수 선언
	var msg string = "Hello variable" // string(문자열) 타입 변수 선언

	a = 20
	msg = "Hello Gopher"
	fmt.Println(a, msg)
}