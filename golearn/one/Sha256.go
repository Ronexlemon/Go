package main

import (
	"fmt"
	"crypto/sha256"
	"crypto"
)
func Sha256(){
	s:="sha256 this string "

	h:=sha256.New()
	h2:= crypto.SHA224.New()
	h2.Write([]byte(s))
	b2:=h2.Sum(nil)

	h.Write([]byte(s))
	bs:= h.Sum(nil)
	
	fmt.Printf("%x\n",bs)
	fmt.Printf("%x\n",b2)
}