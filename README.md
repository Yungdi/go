# The GO Programming Language
## 2 프로그램 구조
### 2.1 이름
- 함수 밖에서 선언되면 같은 패키지 내에서 접근할 수 있다.
- 이름이 대문자로 시작하면 export 돼 자신의 패키지 밖에서 보거나 사용할 수 있다.
- 패키지명 자체는 항상 소문자다.
### 2.2 선언
- 선언의 종류: var, const, type, func
```go
package main

import "fmt"

// 패키지 수준 상수 선언
const boilingF = 212.0

func main() {
    // 함수 지역 변수 선언
    var f = boilingF
    var c = (f - 32) * 5 / 9
    fmt.Printf("boiling point = %gF or %gC\n", f, c)
}
```
- 온도 변환 로직 캡슐화
```go
package main

import "fmt"

func main() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
```