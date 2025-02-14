## 변수

### 변수란?

- 값을 저장하는 메모리 공간을 가리키는 이름

### 변수의 4가지 속성

- 이름 : 변수 이름을 통해 값이 저장된 메모리 공간에 손쉽게 접근할 수 있다
- 값 : 변수가 가리키는 메모리 공간에 저장된 값
- 주소 : 변수가 저장된 메모리 공간의 시작 주소
- 타입 : 변수 값의 형태. 타입을 통해서 사이즈(공간 크기)를 알 수 있다
  - 타입의 사이즈는 고정되어 있다
  - 시작 주소로부터 얼만큼의 크기만큼 읽어야 하는지 타입을 통해 알 수 있다

### 숫자 타입

- uint8 , uint16 , uint32 , uint64
  - unsigned integer
  - 부호 없는 정수 (양수)
  - uint8: 8bit (1byte)
  - uint16: 16bit (2byte)
- int8 , int16 , int32 , int64
  - integer
  - 부호 있는 정수 (음수 + 양수)
- float32 , float64
  - 실수 (floating point : 소수점)
  - float32: 32bit (4byte)
  - float64: 64bit (8byte)

### 그 외 타입

- bool
- string
- 배열
- 슬라이스
- 구조체
- 포인터
- 함수타입
- 맵
- 인터페이스
- 채널

### 별칭(aliasing)

- `byte`: uint8의 별칭 타입
- `rune`: int32의 별칭 타입
  - 문자 1개를 나타내는 타입
  - Go에서는 UTF-8을 문자 인코딩으로 사용
    - UTF-8은 문자 1개를 1byte ~ 3byte로 나타낸다
    - 숫자 타입에서 3byte 존재 X, int32(4byte)에 문자 1개 저장
  - UTF-8로 문자 하나를 나타낼 때 사용
  - rune이 모이면 string이 된다
- `int`: 32bit 컴퓨터에서는 int32 , 64bit 컴퓨터에서는 int64를 의미
- `uint`: 32bit 컴퓨터에서는 uint32 , 64bit 컴퓨터에서는 uint64를 의미

### 변수 선언법

```go
var a int = 10
var a int // 초기값 생략 가능 (초기값을 생략할 경우 기본값이 들어갼다. 정수의 기본값은 0)
var a = 10 // 타입 생략 가능 (숫자 값의 기본 타입은 int) - 타입 생략시 반드시 초기값을 넣어줘야 한다
a := 10 // 선언 대입문
```

### 타입별 기본값

| 타입                                                                                                  | 기본값                                               |
| ----------------------------------------------------------------------------------------------------- | ---------------------------------------------------- |
| 모든 정수 타입 <br> (int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, byte, rune) | 0                                                    |
| 모든 실수 타입 <br> (float32, float64, complex64, complex128)                                         | 0.0                                                  |
| 불리언                                                                                                | false                                                |
| 문자열                                                                                                | "" (빈 문자열)                                       |
| 그 외                                                                                                 | nil (정의되지 않은 메모리 주소를 나타내는 Go 키워드) |

### 타입 변환

- Go는 강타입 언어
- 연산의 각 항목의 타입은 반드시 같아야 한다
- 서로 같은 사이즈의 정수형 타입이라고 해도 서로 타입이 다르면 연산 불가
  - int와 int64는 연산 불가

```go
a := 10 // int 타입 - 64bit 컴퓨터에서는 int64
var b int64 = 20
// 그럼에도 불구하고 a + b 연산 불가
// int와 int64가 서로 같은 사이즈의 정수형 타입이라고 해도 go에서는 서로 다른 타입이기 때문에 연산 불가
// int와 int64는 엄연히 다른 타입
```

```go
// 타입 변환
package main

func main() {
	// 타입 변환
	g := 2 // int
	var h float64 = 2.3

	var i int = int(h) // 2.3 -> 2
	j := float64(g) * h

	var k int64 = 7
	l := g * int(k)

	fmt.Println(g, h, i, j, k, l)
}

```
