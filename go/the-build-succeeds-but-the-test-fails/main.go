package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("I am an error")
	fmt.Printf("v: %v, w: %w", err, err)
}

