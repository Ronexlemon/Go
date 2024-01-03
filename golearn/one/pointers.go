package main

import "fmt"

//passing references to values

//pass by value

func passByValue(n,i int)int{
   return n+i
}


//pass by pointer
func passByreference(n,i *int)int{
	return *n+*i
}

func Pointers(){
	n,i := 3,5
	
	fmt.Println(passByValue(6,4))
	fmt.Println(passByreference(&n,&i))
	//fmt.Println(*0xc0000a61c0);
}
