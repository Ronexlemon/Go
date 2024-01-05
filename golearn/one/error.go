package main

import (
	"fmt"
	"errors"
)


func f1(arg int) (int,error){
	if arg == 42{
		return -1,errors.New("can't work with 42");
	}
	return arg+3,nil
}

//type error

type argError struct{
	arg int
	prob string
}

func (a *argError) Error() string{
	return fmt.Sprintf("%d -%s",a.arg,a.prob)
}

func f2(arg int)(int,error){
	if arg == 42{
		return -1,&argError{arg,"can't work with it"}
	}
	return arg +3 ,nil
}

func Errors(){
	e:=&argError{42,"cant't work with it"}

	fmt.Println(e.Error());
	fmt.Println(f1(42))
	fmt.Println(f2(42))
}