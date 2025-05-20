package main

import (
	"fmt"
)

type Customer struct {
	ID      string
	Balance float64
}

type Store struct {
	m map[string]*Customer
}

// ❌ loop내에서 사용하고있는 c 의 메모리주소를 참조하고있음
func (s *Store) storeCustomers_1(cs []Customer) {
	for _, c := range cs {
		s.m[c.ID] = &c
	}
}

// ❌ current의 가상 메모리 공간을 참조하고있음
func (s *Store) storeCustomer_2(cs []Customer) {
	for _, c := range cs {
		current := c
		s.m[c.ID] = &current
	}
}

// ✅ cs[i] 의 실제 메모리 주소를 참조하고 있음
func (s *Store) storeCustomers_3(cs []Customer) {
	for i := range cs {

		c := &cs[i]
		s.m[c.ID] = c
	}
}

func _22() {
	c := []Customer{
		{ID: "1", Balance: 1.0},
		{ID: "2", Balance: 2.0},
		{ID: "3", Balance: 3.0},
	}

	s := Store{
		m: map[string]*Customer{},
	}

	// 1번방식
	s.storeCustomers_1(c)
	fmt.Println(s)

	// 2번방식
	s.storeCustomer_2(c)
	fmt.Println(s)

	// 3번방식
	s.storeCustomers_3(c)
	fmt.Println(s)

	for _, ccc := range c {
		fmt.Printf("가상 메모리 주소 : %p\n", &ccc) // 가상 메모리 주소 loop 내에서
	}

	for i := range c {
		fmt.Printf("진짜 메모리 주소 : %p\n", &c[i]) // 진짜 메모리 주소
	}
}
