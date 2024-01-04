package main

import (
	"fmt"
)

type base struct{
	num int
}

type container struct{
	base
	str string

}
type describer interface{
	description() string
}

func (b base) description()string{
	return fmt.Sprintf("base with num=%v",b.num);
}

func Structemebending(){
	co:= container{
		base: base{
			num:7,
		},
		str:"yollow",
	}
	fmt.Println("the base is",co.base)
	fmt.Println("the base is",co.base.num)
	fmt.Println("the base is",co.base.description())
	fmt.Println("the base is",co.num)
	fmt.Println("the base is",co.str)

	var d describer = co
	fmt.Println("the base is",d.description())

}