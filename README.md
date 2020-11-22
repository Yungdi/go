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
### 2.3 변수
- var 선언은 특정 타입의 변수를 만들고 이름을 붙인 뒤 초기 값을 설정한다.
```go
//var 이름 타입 = 표현식
var number int = 4
```
- 타입 or 표현식 부분 중 하나는 생략 가능하지만, 둘 다 생략할 수는 없다.  
-> 타입을 생략하는 경우 표현식으로 값을 추론한다. 그래서 값과 타입 동시에 생략 할 수 없어 보인다.  
-> 변수에는 값이 항상 존재한다. 타입에 따라 제로값(초기값)이 다르다.  

```go
var i, j, k int
var b, f, s = true, 2.3, "four"
```
#### 2.3.1 짧은 변수 선언
- 지역 변수를 선언하고 초기화할 수 있는 짧은 변수 선언
```go
// 이름 := 표현식
anim := gif.GIF{LoopCount: nframes}
freq := rand.Float64() * 3.0
t := 0.0

i, j := 0, 1
i, j = j, i
```
#### 2.3.2 포인터
```go
x := 1
p := &x // *int 타입 p는 x를 가리킨다.
fmt.Println(*p) // 1
*p = 2 // x = 2
fmt.Println(x) // 2
```
#### 2.3.3 new 함수
- 변수를 생성하는 또 다른 방법은 내장된 new 함수를 사용하는 것이다.
- new(T) => T 타입의 이름 없는 변수를 만들고 T의 제로 값으로 초기화한 후 *T 타입의 값인 변수의 주소를 반환한다.
```go
p := new(int)
fmt.Println(*p)
```
__같은 함수__
```go
func newInt() *int {
    return new(int)
}

func newInt() *int {
    var dummy int
    return &dummy
}
```
#### 2.3.4 변수의 수명
- 패키지 수준 변수: 프로그램의 전체 실행 시간
- 지역 변수: 동적. 선언문이 실행될 때마다 새 인스턴스가 생성되며, 이 변수는 더 이상 접근할 수 없어서 해당 변수의 저장 공간이 재활용될 때까지 살아 있다.
- GC 는 변수의 저장 공간을 재활용할 수 있는지 여부를 어떻게 알 수 있을까? 해당 변수에 접근할 수 있는 지 여부로 결정된다
```go
var global *int
func f() {
    var x int
    x = 1
    global = &x
}
// x 는 힙에 할당

func g()  {
    y := new(int)
    *y = 1
}
// y 는 스택에 할당 
```
### 2.4 할당
#### 2.4.1 튜플 할당
- 여러 변수를 한 번에 할당할 수 있는 튜플 할당
- 오른쪽의 모든 표현식은 변수가 갱신되기 전에 평가돼 특정 변수가 공교롭게 할당의 양쪽에 모두 나오는 경우에 유용하다
```go
x, y = y, x
a[i], a[j] = a[j], a[i]
```
- 복수의 결과를 반환하는 함수의 호출
```go
f, err = os.Open("foo.txt")
```
#### 2.4.2 할당성
- 묵시적 할당 예시
    1. 함수 호출할 때 파라미터의 할당
    2. return 문의 피연산자에 대응하는 반환 변수
    3. 슬라이스와 같은 복합 타입에 대한 리터럴 표현식
````go
medals := []string{"gold", "silver", "bronze"}
// 위의 문장은 아래와 같
medals[0] = "gold"
medals[1] = "silver"
medals[2] = "bronze"
````
### 2.5 타입 선언
- 변수나 표현식의 타입은 크기, 내부 표현 방식, 수행 가능한 고유 작업, 연관된 메소드와 같은 해당 값의 특성을 정의한다.

## 3 기본 데이터 타입
Go 의 타입  
- 기본
    - 숫자
    - 문자열
    - 불리언
- 결합
    - 배열
    - 구조체
- 참조
    - 포인터
    - 슬라이스
    - 맵
    - 함수
    - 채널
- 인터페이스
### 3.1 정수
- 부호있는 정수: int8, int16, int32(=rune), int64
- 부호없는 정수: uint8(byte), uint16, uint32, uint64

int 와 uint 는 32비트나 64비트지만, 그중 어떤 것일지에 대해서는 가정할 수 없다.
동일한 하드웨어에서라도 다른 컴파일러는 서로 다른 선택을 할 수 있다.
### 3.5 문자열
- 문자열은 불변의 바이트 시퀀스다.  
- 바이트들과 같은 임의의 데이터를 넣을 수 있지만, 보통 사람이 읽을 수 있는 텍스트를 담고 있다.  
- 텍스트 문자열은 통상적으로 유니코드 코드포인트(룬)를 UTF-8로 인코딩한 시퀀스로 해석한다.
- 내장함수: len


## 8 고루틴과 채널
독립적인 작업(goroutine) 간에 값을 전달하지만, 변수는 대부분 단일 작업에 국한되는 모델인 CSP (communicating sequential process)