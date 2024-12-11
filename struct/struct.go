package main

import "fmt"

type User struct {
	Name string
	ID string
	Age int
}

// 구조체를 포함하는 구조체
type VIPUser struct {
	UserInfo User // 내장 타입처럼 포함하는 방식
	VIPLevel int
}

type VIP2User struct {
	User // 포함된 필드(Embedded Field) 방식
	VIPLevel int
}

func main() {
	user := User{"라클", "miracle", 4} 
	vip := VIPUser{
		User{"후추", "Huchu", 2},
		1,
	}
	vip2 := VIP2User{
		User: User{"후추", "Huchu", 2},
		VIPLevel: 2,
	}

	fmt.Printf("유저: %s ID: %s 나이: %d\n", user.Name, user.ID, user.Age)
	fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
	)
	fmt.Printf("VIP2 유저: %s ID: %s 나이: %d VIP 레벨: %d\n",
		vip2.Name,
		vip2.ID,
		vip2.Age,
		vip2.VIPLevel,
	)
}
