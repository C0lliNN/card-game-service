package game

import "fmt"

var (
	ErrInvalidDrawQuantity = fmt.Errorf("the provided draw quantity is not valid")
	ErrDeckNotFound        = fmt.Errorf("no deck was found for the provided id")
)
