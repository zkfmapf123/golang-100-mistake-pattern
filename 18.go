package main

/*
❌ copy 안됨
src, dst 모두 동일한 길이를 가지고 있어야 함
*/
func badCopy() []int {
	src := []int{0, 1, 2}
	var dst []int // nil slice (메모라 할당 x)

	copy(dst, src)
	return dst
}

/*
✅ copy 됨
src, dst 동일한 길이를 가지게 함
*/
func goodCopy() []int {
	src := []int{0, 1, 2}
	dst := make([]int, len(src))

	copy(dst, src)
	return dst
}
