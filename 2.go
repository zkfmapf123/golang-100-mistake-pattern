package main

func BadIfPattern(num int) {
	if num > 10 {
		// ...
		if num > 20 {
			// ...
			if num > 30 {
				// ...
			} else {
				// ...
			}
		} else {
			// ...
		}
	} else {
		//...
	}
}

// 최대한 else 쓰지 마라 (중첩 쓰지말자...)
func GoodIfPattern(num int) {
	if num > 10 {
		// return ...
	}

	if num > 20 {
		// return ...
	}

	if num > 30 {
		// return ...
	}
}
