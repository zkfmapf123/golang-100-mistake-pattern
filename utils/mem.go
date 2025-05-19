package utils

import (
	"fmt"
	"runtime"
)

func PrintMemory(msg string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("%s : ", msg)
	fmt.Printf("Alloc = %v MiB\t", bToMb(m.Alloc))
	fmt.Printf("TotalAlloc = %v MiB\t", bToMb(m.TotalAlloc))
	fmt.Printf("Sys = %v MiB\t", bToMb(m.Sys))
	fmt.Printf("NumGC = %v\t", m.NumGC)
	fmt.Println("\n")
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
