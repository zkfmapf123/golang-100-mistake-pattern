package main

/*
❌
golang에서의 반복문은 runtim시, 맵 항목에 값이 생성된다면
값이 생성될 수도 있고, 안될 수 도 있다. (랜덤) -> 반복회차에 따라 달라질 수 있음

1 시도 : map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true]
2 시도 : map[0:true 1:false 2:true 10:true 12:true 20:true 22:true 30:true 32:true 40:true]
3 시도 : map[0:true 1:false 2:true 10:true 12:true 20:true 22:true]
*/
func _23_loopUpdate() map[int]bool {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	for k, v := range m {
		if v {
			m[10+k] = true
		}
	}

	return m
}

/*
✅ 복사본을 만들어서 해결하자... (copyMap)
>> map[0:true 1:false 2:true 10:true 12:true]
*/
func _23_goodLoopUpdate() map[int]bool {
	m := map[int]bool{
		0: true,
		1: false,
		2: true,
	}

	copyMap := func(src map[int]bool) map[int]bool {
		dst := make(map[int]bool, len(src))
		for k, v := range src {
			dst[k] = v
		}

		return dst
	}

	dm := copyMap(m)

	for k, v := range m {
		if v {
			dm[10+k] = true
			continue
		}

		dm[k] = v
	}

	return dm
}
