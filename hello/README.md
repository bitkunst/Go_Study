## Go

### package

- 패키지명은 임의로 정할 수 있다.
- 패키지명이 `main`인 경우, 프로그램 시작점을 포함하는 패키지라는 의미
  - Go로 만든 프로그램은 main 패키지 1개와 여러 개의 다른 패키지들(main이 아닌 패키지)로 구성된다.

### import "packageName"

- 특정 기능을 가지고 있는 패키지를 가져와서 사용하고 싶을 때 `import` 구문을 통해 패키지를 가져와서 사용한다.

### func main

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello Gopher!")
}
```

- `main 함수`는 프로그램 시작점을 나타낸다.
- `package main`: 프로그램 시작점을 포함하는 패키지 (main 함수를 가지고 있는 패키지)
