package main // Go의 모든 코드는 pacakage로 시작한다 (해당 코드가 속한 패키지)
// Go에서 모든 코드는 특정 패키지 안에 들어가 있어야 한다
// package명은 임의로 정할 수 있다
// package명이 main인 경우, 프로그램 시작점을 포함하는 패키지라는 의미
// main 패키지 하나와 여러 개의 다른 패키지로 구성된다

import "fmt" // fmt 패키지를 가져온다

// main 함수는 프로그램 시작점을 나타낸다
// main 함수를 가지고 있는 패키지가 package main
func main() {
	fmt.Println("Hello, World!") // fmt 패키지 안에 있는 Println 함수를 사용
}