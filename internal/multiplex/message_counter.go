package multiplex

type MessageCounter[T any] struct {
	Count   int
	Last    bool
	Message T
}
