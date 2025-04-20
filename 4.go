package main

// Sort Interface
type Impl interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func IsSorted(data Impl) bool {
}
