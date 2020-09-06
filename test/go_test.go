package test

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

/*
* 字符是放常量池, 对象比较的是地址
*/
func TestError(t *testing.T) {
	err1 := errors.New("test")
	err2 := errors.New("test")
	if err1 == err2 {
		fmt.Print("equal")
	}else {
		fmt.Print("not equal")
	}
    fmt.Println("")
	s1 := "test"
	s2 := "test"
	if s1 == s2 {
		fmt.Print("equal")
	}else {
		fmt.Print("not equal")
	}
	fmt.Println("")
}

func TestTime(t2 *testing.T) {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	fmt.Print(testReturn())
}

func testReturn()(res string){
	res = "hello"
	return "world"
}
