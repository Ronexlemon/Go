package main

import (
	"fmt"
	"unicode/utf8"
)

const st = "สวัสดี"


func returnstringLength()int{
	return len(st);
}

//indexing into a string to produce raw byte at each index
func countRunesInString(){

	for idx, runeValue := range st {
        fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }
	fmt.Println(utf8.RuneCountInString(st));
}
func indexing(){
	
	for i:=0;i < len(st);i++{
		fmt.Printf("%x \n ",st[i])

	}
}

//using utf8 to count the number of runes in a string


func StringRune(){
	
	fmt.Println(returnstringLength())
	indexing()
	countRunesInString()
}