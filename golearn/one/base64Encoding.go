package main

import (
	"fmt"
	b64 "encoding/base64"
)

func Base64Encoding(){
	data := "yollow124653brbh you"
	standardencoder:= b64.StdEncoding.EncodeToString([]byte(data))

	fmt.Println(standardencoder)
	//decoding
	standarddecoding,_:= b64.StdEncoding.DecodeString(standardencoder)
	fmt.Println(string(standarddecoding))

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}