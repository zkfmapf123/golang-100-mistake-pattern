package main

import "fmt"

// 기본적인 GetKeys
func GetKeys(m map[string]int) []string {
	var keys []string

	for k := range m {
		keys = append(keys, k)
	}

	return keys
}

// 여러 타입을 받아야 한다면?
func GetKeysOtherTypes(m any) ([]any, error) {
	switch t := m.(type) {

	default:
		return nil, fmt.Errorf("unknown type : %T", t)

	case map[string]int:
		var keys []any
		for k := range t {
			keys = append(keys, k)
		}
		return keys, nil

	case map[int]string:
		// ...
	}

	return nil, nil
}

// ---------------------------- 제네릭을 사용한다면? ✅ ----------------------------
// Type을 제한해야 한다면

// int 와 string 만 강제한다
type ComparableType interface {
	~int | ~string
}

func GetKeysUseGeneric[K ComparableType, V any](m map[K]V) []K {
	var keys []K
	for k := range m {
		keys = append(keys, k)
	}

	return keys
}
