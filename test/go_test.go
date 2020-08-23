package test

import (
	"errors"
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	err1 := errors.New("test")
	err2 := errors.New("test")
	if err1 == err2 {
		fmt.Print("equal")
	}else {
		fmt.Print("not equal")
	}

	s1 := "test"
	s2 := "test"
	if s1 == s2 {
		fmt.Print("equal")
	}else {
		fmt.Print("not equal")
	}
}
