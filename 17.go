package main

import (
	"fmt"
)

func generateId(id string) string {
	return fmt.Sprintf("%s-%s", id, id)
}

func handleOperations(id string) []any {
	op := generateId(id)

	// nil slice를 판별하는 최고의 함수
	if len(op) == 0 {
		return nil
	}

	return []any{} // empty slice
}
