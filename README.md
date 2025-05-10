# Golang 100 가지 실수패턴과 솔루션

- ❌ Bad
- ✅ Good

- [1. 의도하지 않은 변수가림을 조심하라](./1.go)
- [2. 중첩된 if 문을 피하라](./2.go)
- [3. init 함수를 피하라 / 그럴거면 전역변수 + struct + 명시적으로 호출하라](./3.go)
- [4. 인터페이스]()
    - [제공자 측에 인터페이스를 두지마라 - 사용자 측에 둬라](./4.interface)
    - [any를 막쓰지 마라](./4.interface)
- [5. 제네릭이 필요한 시점을 파악해라 / 근데 너무 많이쓰면 더 어려워짐...](./5.go)
- [11. 함수형 패턴을 사용하라](./11.함수현_패턴/)
- 12. 프로젝트 구성

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
- 13. linter 활용

<a href="https://golangci-lint.run/welcome/install/#local-installation"> golangci-lint </a>

```sh
    golangci-lint
```
