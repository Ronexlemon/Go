package main

import (
	"fmt"
	"bufio"
	"os"
)

func checkerror(e error){
	if e !=nil{
		panic(e)
	}
	
}

func WriteFiles(){
	d1:=[]byte("//readingAFile")
	err:=os.WriteFile("./todo.txt",d1,0644)
	checkerror(err)
	f,err:=os.Create("./todo2.txt")
	checkerror(err)

	defer f.Close()

	d2:=[]byte{115,111,109,101,10}
	n2,err:= f.Write(d2)
	checkerror(err)
	fmt.Printf("wrote %d bytes\n",n2)


	n3,err:=f.WriteString("//writes string\n")
	checkerror(err)
	fmt.Printf("wrote %d bytes\n",n3)

	f.Sync()

	w:= bufio.NewWriter(f)
	n4,err:= w.WriteString("buffered \n")
	checkerror(err)
	fmt.Printf("wrote %d bytes\n",n4)

	w.Flush()

}
