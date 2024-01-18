package errorhandle

import (
	"fmt"
	"errors"
)

func ErrorHandle(name string)(string,error){
	//if no name was given return witha message
	if name ==""{
		return "", errors.New("empty name")

	}
	return fmt.Sprintf("Hi, %v. Welcome",name),nil
}