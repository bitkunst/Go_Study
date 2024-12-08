package main

import "fmt"

// main 함수:  프로그램 시작점
// Go에서 모든 프로그램은 main에서 시작해서 main에서 끝난다
func main() {
	var a int = 10
	var b int = 20
	var f float64 = 23465714325.6617

	var c int
	var d int

	fmt.Print("a:", a, "b:", b) // 입력값들이 빈칸 없이 출력
	fmt.Println("a:", a, "b:", b, "f:", f) // 입력값들 사이사이에 빈칸 하나씩 넣어준다 + 자동 개행
	fmt.Printf("a: %d b: %d f: %f \n", a, b, f) // 서식 문자(formatter) 자리에 파라미터로 전달한 값이 들어가서 출력

	n, err := fmt.Scanln(&c, &d) // &는 메모리 주소값을 나타낸다 (&c == c변수의 메모리 주소값)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(n, c, d)
	}
}