package main

import (
	"fmt"
	"slices"
)

func unInitialisedSlice(){
	var s []string
	
	fmt.Println("uninit",s,s ==nil,len(s)==0)
}
func emptysliceWithnonZero(){
	s:=make([]string,3)
	fmt.Println("emp",s,"len",len(s),"cap",cap(s))
	s[0]="a"
	s[1]="b"
	s[2]="c"
	
	fmt.Println("After",s,"len",len(s),"cap",cap(s))

}
func appendslice()[]string{
	s:= make([]string,3)
	s=append(s, "d")
	s=append(s, "e","f","g")
	
	return s
}

func copySlice(){
	s:= appendslice()
	c:=make([]string,len(s))
	copy(c,s)
	fmt.Println("copy s to c",c)
	
}
func sliceASlice(){
s:= appendslice()
sli:= s[2:5]
fmt.Println("slice from index 2 to 5 excluding index 5",sli)
sli2:= s[:5]
fmt.Println("slice from index 0 to 5 excluding index 5",sli2)
sli3:=s[2:5]
fmt.Println("slice from index 2 to last index",sli3)
}

func straightLineSlices()[]string{
	t:=[]string{"a","b","c","v"}
    return t
}
func checkEqualSlices(){
	s:=straightLineSlices()
	s1:=straightLineSlices()
	if slices.Equal(s,s1){
		fmt.Println("s ==s1")
	}
}
func twoDSlice(){
	twoD:=make([][]int,3)

	for i:=0;i<3;i++{
		innerLen:=i+1
        twoD[i] = make([]int, innerLen)
		for j:=0;j <innerLen;j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
func Slice(){
	unInitialisedSlice()
	emptysliceWithnonZero()
	copySlice()
	sliceASlice()
	straightLineSlices()
	checkEqualSlices()
	twoDSlice()
}