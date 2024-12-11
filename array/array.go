package main

import (
	"fmt"
)

func main() {
	// int 타입 요소를 5개 갖는 배열 nums 할당
	var nums [5]int = [5]int{1:10, 3:30} 
	// 인덱스가 1인 요소값을 10으로, 3인 요소값을 30으로 초기화
	// 나머지는 int 타입의 기본값인 0으로 초기화

	// ...를 사용해 배열 요소 개수 생략 가능
	x := [...]int{10, 20, 30} // 이때 배열 요소 개수는 초기화되는 요소 개수와 같다
	xLength := len(x) // 내장함수 len()으로 배열 길이를 알 수 있다

	fmt.Println(nums)
	fmt.Println(x, xLength)

	var t [5]float64 = [5]float64{24.0, 25.9, 27.8, 26.9, 26.2}
	// for 반복문에서 range 키워드를 이용하여 배열 요소 순회 가능
	// range는 배열의 각 요소를 순회하면서 인덱스와 요소값 두 값을 반환
	for i, v := range t {
		// i는 인덱스, v는 원소값
		fmt.Println(i, v)
	}

	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{500, 400, 300, 200, 100}
	fmt.Println(a == b) // false

	for i, v := range a {
		fmt.Printf("a[%d] = %d\n", i, v)
	}

	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}

	b = a
	fmt.Println(a == b) // true
	for i, v := range b {
		fmt.Printf("b[%d] = %d\n", i, v)
	}
}