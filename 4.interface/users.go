package main

// 사용자 측에 인터페이스를 두어라...
type CustomerStorage interface {
	Get(any) error // ❌
	Set(any) error // ❌

	GetUserId(int) error          // ✅
	GetCusomterName(string) error // ✅
	SetUserId(int) error          // ✅
	SetUserName(string) error     // ✅

	Marshal(v any) // ✅ 보통 marshal 하거나, 애매한 매개변수는 any로 주자
}
