# Golang 100 가지 실수패턴과 솔루션

- ❌ Bad
- ✅ Good

### [1. 의도하지 않은 변수가림을 조심하라](./1.go)
### [2. 중첩된 if 문을 피하라](./2.go)
### [3. init 함수를 피하라 / 그럴거면 전역변수 + struct + 명시적으로 호출하라](./3.go)
### [4. 인터페이스]()
    - [제공자 측에 인터페이스를 두지마라 - 사용자 측에 둬라](./4.interface)
    - [any를 막쓰지 마라](./4.interface)
### [5. 제네릭이 필요한 시점을 파악해라 / 근데 너무 많이쓰면 더 어려워짐...](./5.go)
### [11. 함수형 패턴을 사용하라](./11.함수현_패턴/)
### 12. 프로젝트 구성

```sh
|- cmd                  -- 메인 소스 파일
|- internal             -- private code
|- pkg                  -- public code
|- test                 -- 단위테스트가 아닌, 공용 API 테스트 및 통합테스트는 해당 폴더에 위치
|- configs              -- 설정 파일
|- docs
|- examples
|- api                  -- swagger 및 protocol buffer를 둔다
|- web                  -- 정적파일
|- build                -- CI 파일 위치
|- scripts
|- vendor               -- 의존성 파일
```
### 13. linter 활용

<a href="https://golangci-lint.run/welcome/install/#local-installation"> golangci-lint </a>

```sh
    golangci-lint
```

### 14. 부동소스점을 정확하게 이해해라 (...)

### [15. 슬라이스의 길이와 용량을 정확하게 이해하라](./15.go)

- 슬라이스의 개념을 잘 모르면 -> 슬라이스 관련 성능이 떨어질 수 있음
- capa는 최대용량임 -> 이걸 넘어가면 <b>cap = cap*2</b> 로 증가함
- <b>대신 capa가 1024 이상이면 *2 가 아닌, 25% 씩 증가함</b>

```go

// len 3, cap 6
s := make([]int, 3,6)

s[3] = 10 // range error ❌
s = append(s, 10) // ✅

```

### [16. nil과 빈 슬라이스를 혼동하지 마라](./16.go)

- 슬라이스의 길이가 0이면 <b>empty slice</b> 다
- 슬라이스가 nil과 같으면 <b>nil slice</b> 다

```go
    []string                // nil slice
    []string(nil)           // nil slice

    []string{}              // empty slice -> 할당을 해줘야 함
    make([]string,0)        // empty slice -> 할당을 해줘야 함
```

### [17. 슬라이스가 비었는지 제대로 확인하라](./17.go)

- nil slice, empty slice 를 판별하는 함수를 작성하라
- nil , empty slice를 판별하는 피하라고 권장함
- map 자체를 검사하려면 <b>nil 여부를 확인하지 말고 길이를 확인하자</b>

### [18. 부정확한 슬라이스 복제 관련 실수](./18.go)

- copy 는 src, dst 모두 동일한 길이를 가져야 함
- copy를 사용하지 않는다면, append를 사용해라

### [19. 슬라이스와 메모리 누수 관련 실수](./19.go)

```go

// Bad Code ❌
a := []int{...100,000,000} // len 100,000,000, cap 100,000,000
b : =a[:5] // len 5, cap 100,000,000

// Good Code ✅
a := []int{...100,000,000} // len 100,000,000, cap 100,000,000
b := make([]int, num)

copy(b, a)

b // len 5, cap 5

```

### [20. 값을 올바르게 비교하는 방법](./20.go)

- 서로 값을 비교할때는 reflect를 이용해라
- reflect는 메타프로그래밍으로서, app 내부구조와 동작을 변경함
- reflect.DeepEqual은 재귀적으로 탐색해서 동등관계를 파악함

```go

// ❌ bad
type customer struct {
    id string
}

c1 := cusomter{id : "x"}
c2 := cusomter{id : "y"}

fmt.Println(c1 == c2) // false

```

```go

// ❌ bad (compile error)
type customer struct {
    id string
    op []float64
}

c1 := cusomter{id : "x", op : []float64{1.0}}
c2 := cusomter{id : "y", op : []float64{1.0}}

fmt.Println(c1 == c2) // compile error

```

### [21. range 안에서의 값은 복제본이다](./21.go)

- range loop 에서 값에 해당하는 원소는 항상 <b>복제본</b>이다.
- slice 참조는 값이 바뀐다.
- 배열은 값이 바뀌지 않는다
- 배열을 range 내에서 값을 바꾸려면 <b> 배열 포인터를 사용해라 ...</b>

```go
    for _, v := range []int{1,2,3,4,5} {
        // v <- 복제본임...
    }
```

### [22. loop 내에서 반복문의 참조값을 파악해라](./22.go)

- 루프변수 c는 각 반복마다 새로운 메모리 공간을 할당하는 것이 아닌, 같은 메모리 공간을 재사용 한다.
- 만약 cs 실제값들의 메모리 주소를 참조해야 한다면 index 값으로 직접 지정해야 한다.
- 위에서 말했듯이, loop 변수는 각 반복마다 새로운 = 같은 메모리공간을 사용한다. 그래서 값 자체의 대한 변조 및 참조는 index를 사용해야한다.

```go

// 가짜 메모리 주소
for _, c := range cs {
    &c // <- 0x00014
}

// 진짜 메모리 주소
for i := range cs {
    &cs[i] // <- 0x00014
}


```
