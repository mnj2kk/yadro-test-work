package structures

type Pair[T, U any] struct {
	First  T
	Second U
}

type Result []Pair[string, int]
