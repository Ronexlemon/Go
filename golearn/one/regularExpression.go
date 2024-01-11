package main
import (
	//"bytes"
	"fmt"
	"regexp"
)

func RegularExpression(){
	match,_:=regexp.MatchString("([a-z])","peach")
	match2,_:=regexp.Compile(`^\+[1-9]\d{1,14}$`)
	re,_:=regexp.Compile("p([a-z]+)ch")
	fmt.Println("is: ",match) 
	fmt.Println(match2.MatchString("+254701707772"))
	fmt.Println(re.MatchString("peach"))
	fmt.Println("idx: ",re.FindStringIndex("peach punch"))
	fmt.Println(re.FindStringSubmatch("peach punch"))
	fmt.Println(re.FindAllString("peach punc",1))


}