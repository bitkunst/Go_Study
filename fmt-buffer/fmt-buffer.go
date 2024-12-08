package main

import (
	"bufio"
	"fmt"
	"os"
)

// 표준 입력 실패시 입력 버퍼를 비우는 방법
func main() {
	// 표준 입력 버퍼
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	n, err := fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n') // 표준 입력 버퍼에서 개행문자('\n')가 나올 때까지 읽어오라는 뜻 -> \n이 나올 때까지 버퍼가 지워짐
	} else {
		fmt.Println(n, a, b)
	}

	n, err = fmt.Scanln(&a, &b)
	if err != nil {
		fmt.Println(err)
		stdin.ReadString('\n')
	} else {
		fmt.Println(n, a, b)
	}
}