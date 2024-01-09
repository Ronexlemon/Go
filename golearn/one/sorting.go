package main
import (
	"fmt"
	"slices"

)

func returnIntSlice()[]int{
	sl:=[]int{1,2,30,4,50,6}
	return sl

}
func returnStringSlice()[]string{
	st:=[]string{"f","g","k","h","a"}
	return st
}

func Sort(){
	intSlice:= returnIntSlice()
	stringSlice:= returnStringSlice()
	slices.Sort(intSlice)
	slices.Sort(stringSlice)
	fmt.Println("sorted int slice",intSlice)
	fmt.Println("sorted string slice",stringSlice)
	fmt.Println(slices.IsSorted(intSlice))

}


