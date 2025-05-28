package main

import (
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"
)

var workerPool = runtime.GOMAXPROCS(0)

/*
	워크로드 타입에 따른 동시성 구성
	- Cpu Bounded
		- 병합정렬 알고리즘
	- IO Bounded
		- REST 호출
		- DB CRUD
	- Virutal Memory (별로 중요하지 않음)
		- 메모리 중심

	워커풀링 시, 워크로드 타입에 따른 개수
	- CPU Bounded -> 최적의 고루틴 수는 가용스레드의 개수와 가까움 (runtime.GOMAXPROCS(0))
	- IO Bounded -> 외부시스템에 따라 풀의크기가 결정됨

	결국...
	- 벤치마크를 통해 상황에 맞게 구성해야 함
*/

func read() []byte {
	src := fmt.Sprint("jobs/file.txt")
	b, err := os.ReadFile(src)
	if err != nil {
		return nil
	}
	return b
}

func write(filename string, b []byte) {
	dst := fmt.Sprintf("dist/%s.txt", filename)
	err := os.WriteFile(dst, b, 0o644)
	if err != nil {
		return
	}
}

// 그냥 task
func task() {
	b := read()
	write("1", b)
}

/*
✅ 순차 쓰기
112.50ms
*/
func taskLoop() {
	start := time.Now()

	for i := 0; i < workerPool; i++ {
		b := read()
		write(fmt.Sprintf("%d", i), b)
	}

	elapsed := time.Since(start)
	fmt.Printf("순차 for-loop 소요시간: %s\n", elapsed)
}

/*
✅ WaitGroup을 사용해서 쓰기
52.09ms
*/
func taskWaitGroup() {
	start := time.Now()

	wg := sync.WaitGroup{}

	for i := 0; i < workerPool; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			b := read()
			write(fmt.Sprintf("%d", i), b)
		}(i)
	}

	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("WaitGroup 소요시간: %s\n", elapsed)
}

func processWithWorkerPool(workerCount int, filenames []string) error {
	// 작업 시작 시간 측정
	start := time.Now()

	// 작업 채널과 결과 채널 생성
	jobs := make(chan string, workerCount)
	results := make(chan error, workerCount)

	// WaitGroup 생성
	var wg sync.WaitGroup

	// 워커 시작
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for filename := range jobs {
				// 파일 읽기
				b := read()
				if b == nil {
					results <- fmt.Errorf("파일 읽기 실패: %s", filename)
					continue
				}

				// 파일 쓰기
				write(filename, b)
				results <- nil
			}
		}()
	}

	// 작업 분배
	go func() {
		for _, filename := range filenames {
			jobs <- filename
		}
		close(jobs)
	}()

	// 워커 완료 대기
	go func() {
		wg.Wait()
		close(results)
	}()

	// 에러 체크
	for err := range results {
		if err != nil {
			return err
		}
	}

	// 소요 시간 출력
	elapsed := time.Since(start)
	fmt.Printf("Worker Pool 처리 소요시간: %s\n", elapsed)

	return nil
}
