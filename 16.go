package main

/*
✅ 슬라이스 리턴함수는 nil slice를 리턴해야 한다 (할당하지 않아도 되니까...)

장점
- 메모리 효율성 (실제로 메모리를 할당하지 않는다) <=> empty slice는 메모리를 할당함
- nil slice는 null 로 직렬화됨 <=> empty slice는 []로 직렬화됨
- nil slice는 "데이터가 없음" 을 명확히 표현 <=> empty slice는 "빈 배열이 있음"
*/
func GetUserNilSlice() []string {
	var s []string

	if s != nil {
		return nil
	}

	return s
}

/*
❌ empty slice
*/
func GetUserEmptySlice() []string {
	return []string{} // 메모리 할당
}
