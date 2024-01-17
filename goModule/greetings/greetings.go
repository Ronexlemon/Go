package greetings

import (
	"fmt"
)

func HelloReturnName(name string)string{
	message:= fmt.Sprintf("Hi, %v. Welcome!,",name)
	return message
}