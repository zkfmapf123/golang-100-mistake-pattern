package main

import (
	"fmt"
	"log"
)

var (
	config map[string]string
	db     *Database
)

type Database struct {
	connection string
}

func NewDatabase(conn string) *Database {
	return &Database{connection: conn}
}

// /////////////////////////////////////////////// Bad Pattern - init /////////////////////////////////////////////////
func init() {
	// 1. 복잡한 로직을 init에 넣는 것은 X
	config = make(map[string]string)
	config["host"] = "localhost"
	config["port"] = "5432"

	// 2. 에러 처리가 어려움
	db = NewDatabase("connection string")

	// 3. 테스트가 어려움
	fmt.Println("데이터베이스 초기화 완료")
}

// /////////////////////////////////////////////// Good Pattern - init /////////////////////////////////////////////////
func setupConfig() error {
	config = make(map[string]string)
	config["host"] = "localhost"
	config["port"] = "5432"
	return nil
}

func setupDatabase() error {
	var err error
	db, err = connectDatabase()
	if err != nil {
		return err
	}
	return nil
}

func connectDatabase() (*Database, error) {
	return NewDatabase("connection string"), nil
}

// 초기화 로직을 명시적으로 호출
func GoodInitPattern() {
	if err := setupConfig(); err != nil {
		log.Fatal("설정 초기화 실패:", err)
	}

	if err := setupDatabase(); err != nil {
		log.Fatal("데이터베이스 연결 실패:", err)
	}

	fmt.Println("애플리케이션 시작")
}

/*
init 함수 사용 시 주의사항:

1. init 함수는 프로그램 실행 시 자동으로 호출되며, 제어가 어렵습니다.
2. init 함수에서 발생하는 에러는 처리하기 어렵습니다.
3. init 함수는 테스트하기 어렵습니다.
4. init 함수의 실행 순서는 패키지 의존성에 따라 달라질 수 있습니다.

Best Practices:

1. 명시적인 초기화 함수를 사용하세요:
   - setup(), initialize() 등의 함수를 만들어 필요한 곳에서 명시적으로 호출
   - 에러 처리가 가능
   - 테스트가 용이

2. 의존성 주입을 사용하세요:
   - 생성자 함수를 통해 의존성을 주입
   - 테스트 시 mock 객체로 교체 가능

3. 전역 상태를 최소화하세요:
   - 전역 변수 대신 구조체를 사용
   - 필요한 곳에서 인스턴스를 생성

4. 초기화 로직을 main 함수에서 관리하세요:
   - 명확한 초기화 순서
   - 에러 처리 가능
   - 로깅 가능
*/
