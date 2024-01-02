package main

import ("fmt")


func fact(n int) int{
	if n==0{
		return 1
	}
	return n * fact(n-1)
}

func fib(n int)int{
	var fibb  func(n int)int
	fibb =func(n int)int{
		if n <2 {
			return n
		}
		return fibb(n -1) + fibb(n-2)
	}
	return fibb(n)
}

func Recursion(n int){
	fmt.Println(fact(n))
	 fmt.Println(fib(n))


}