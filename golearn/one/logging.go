package main

import (
	"fmt"
	//"bytes"
	"bytes"
	"log"
	"os"
	"log/slog"
)

func Logging(){
	log.Println("Standard logger")
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout,"my:",log.LstdFlags)
	myLog.Println("from my log")

	myLog.SetPrefix("ohmy:")
	myLog.Println("from mylog")

	var buf bytes.Buffer
	buflog:= log.New(&buf,"buf:",log.LstdFlags)

	buflog.Println("hello")
	fmt.Print("from buflog:",buf.String())


	jsonHandler:= slog.NewJSONHandler(os.Stderr,nil)
	myslog:= slog.New(jsonHandler)

	myslog.Info("hello there")

	myslog.Info("hello once again","key","val","age",25)

}