package main

import (
	"fmt"
	"net"
	//"net"
	"net/url"
)

func URLParsing(){
	s:= "postgress://user:pass@host.com:4040/path?k=v#f"
	u,err:= url.Parse(s)

	if err !=nil{
		panic(err)
	}
	fmt.Println(u.Scheme)

	fmt.Println(u.User)
	fmt.Println(u.User.Username())

	p,_:=u.User.Password()
	fmt.Println(p)
	fmt.Println(u.Host)
	host,port,_:= net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m,_:= url.ParseQuery(u.RawQuery)
	fmt.Println(m)
}