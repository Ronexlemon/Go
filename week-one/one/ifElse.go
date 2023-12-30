package main

import (
	"fmt"
)


func basicIfAndElse(num int)string{
	if num%2 ==0{
		return "even"
	}else{
		return "odd"
	}

	

}

func logicOperator(num int)string{
	//logic operator
	if num % 2 ==0 && num ==2{
		return "its a two"
	}
	return "other number"
}


func precedeCondition(i int){
	if num:=i; num <0{
		fmt.Println("negative number")

	}else if num <10{
		fmt.Println("the number has one digit")
	}else{
		fmt.Println("has multiple digits")
	}
}

func IfElse(){
	fmt.Println(basicIfAndElse(4))
	fmt.Println(logicOperator(8))
	precedeCondition(9)
}