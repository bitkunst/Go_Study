## Struct

- 구조체(Struct)는 여러 필드들을 묶은 타입
- `type` 키워드를 사용해 사용자 정의 타입을 정의
- 타입명의 첫번째 글자가 대문자이면 패키지 외부로 공개되는 타입
- 각 필드는 필드명과 타입을 적는다

```go
type 타입명 struct {
    필드명 타입
    필드명 타입
}

type Student struct {
    Name string
    Class int
    No int
    Score float64
}
```

### 구조체 변수 초기화

#### 1. 초기값 생략

- 초기값을 생략하면 모든 필드가 기본값으로 초기화된다

```go
type House struct {
    Address string
    Size int
    Price float64
    Type string
}

var house House
// house의 필드값은 다음과 같이 초기화
/**
 * Address: ""
 * Size: 0
 * Price: 0.0
 * Type: ""
*/
```

#### 2. 모든 필드 초기화

- 모든 필드값을 중괄호 `{}` 사이에 넣어서 초기화한다
- 모든 필드가 순서대로 초기화된다

```go
var house House = House{ "서울시 강남구", 28, 19.5, "아파트" }
/**
 * Address: "서울시 강남구"
 * Size: 28
 * Price: 19.5
 * Type: "아파트"
*/
```

#### 3. 일부 필드 초기화

- 일부 필드값만 초기화할 때는 `"필드명: 필드값"` 형식으로 초기화한다
- 초기화되지 않은 나머지 변수에는 기본값이 할당된다

```go
var house House = House{ Size: 28, Type: "아파트" }
/**
 * Address: ""
 * Size: 28
 * Price: 0.0
 * Type: "아파트"
*/
```

### 구조체를 포함하는 구조체

- 구조체의 필드로 다른 구조체를 포함할 수 있다
- 일반적인 `내장 타입처럼 포함하는 방법`과 `포함된 필드(Embedded Field) 방식`이 있다
  - 구조체 안에 포함된 다른 구조체의 필드명을 생략하는 경우를 `포함된 필드`라고 부른다

```go
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
// 구조체에서 다른 구조체를 필드로 포함할 때 필드명을 생략하면 .을 한번만 찍어 접근할 수 있다

func main() {
	user := User{"라클", "miracle", 4}
	vip := VIPUser{
		User{"후추", "Huchu", 2},
		10,
	}
    vip2 := VIP2User{
		User: User{"후추", "Huchu", 2},
		VIPLevel: 2,
	}

    // Name 필드는 vip 변수의 UserInfo 필드 안에 속하기 때문에 vip.UserInfo.Name으로 접근해야 한다
    fmt.Printf("VIP 유저: %s ID: %s 나이: %d VIP 레벨: %d\n",
		vip.UserInfo.Name,
		vip.UserInfo.ID,
		vip.UserInfo.Age,
		vip.VIPLevel,
	)
    // Embedded Field 방식으로 정의하면 보다 편하게 접근 가능
	fmt.Printf("VIP2 유저: %s ID: %s 나이: %d VIP 레벨: %d\n",
		vip2.Name,
		vip2.ID,
		vip2.Age,
		vip2.VIPLevel,
	)

}
```

### 복사

- 구조체 변수값을 다른 구조체에 대입하면 모든 필드값이 복사된다
- Go 내부에서는 필드 각각이 아닌 구조체 전체를 한번에 복사한다
  - 대입 연산자가 우변 값을 좌변 메모리 공간에 복사할 때 '복사되는 크기'는 '타입 크기'와 같다
  - 구조체 크기는 모든 필드를 포함하므로 구조체 전체 필드가 복사되는 것이다

```go
type Student struct {
    Age int
    No int
    Score int
}

func main() {
    var student = Student{19, 23, 98.2}

    // student 구조체의 모든 필드가 student2로 복사된다
    student2 := student
    fmt.Println(student2) // {19 23 98.2}
    fmt.Println(student == student2) // true
}
```
