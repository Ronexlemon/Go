package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LineFilters(){
	scanner:= bufio.NewScanner(os.Stdin)
	fmt.Println("input value to scan")
	for scanner.Scan(){
		if(scanner.Text() == "me"){
			break
		}
		ucl:= strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	if err:= scanner.Err();err !=nil{
		fmt.Fprintln(os.Stderr,"error",err)
		os.Exit(1)
}
}