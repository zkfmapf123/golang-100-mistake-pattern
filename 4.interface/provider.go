package main

type Customer struct{}

// 제공자 측에 인터페이스를 두지 마라 !!!
// type CustomerStorage interface {
// 	StoreCustomer(Customer) error
// 	GetCustomer(id string) error
// 	UpdateCustomer(Customer) error
// 	GetAllCustomer() ([]Customer, error)
// }

/*
- any를 너무쓰게 되면 함수 시그니처만 보고 판단하기 어려워진다.
- 함수 자체의 대한 기능이나 의도가 명확하게 보이지 않음
- Solution) 각 타입마다 함수를 각각 만들어아 한다. ********
*/
type (
	A struct{}
	B struct{}
	C struct{}
)

// ❌ Bad
func (c *C) Get(id string) (any, error) {
	return "", nil
}

// ❌ Bad
func (c *C) Set(id string, v any) error {
	return nil
}

// ✅ Good
func (c *C) GetUserId(id int) (int, error) {
	return 0, nil
}

// ✅ Good
func (c *C) GetCustomerName(name string) (string, error) {
	return "", nil
}

// ✅ Good
func (c *C) SetUserId(id int) error {
	return nil
}

// ✅ Good
func (c *C) SetCustomerName(name string) error {
	return nil
}
