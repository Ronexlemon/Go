package randomGreeting

import (
	"fmt"
	"math/rand"
	"errors"
)

func randomFormat()string{
	formats:=[]string{"Hi, %v. Welcome!","Great to see you, %v","Hail, %v! well met!",}
	return formats[rand.Intn(len(formats))]
}
func RandomGreeting(name string)(string,error){
	if name ==""{
		return "", errors.New("empty name")
	}

	return fmt.Sprintf(randomFormat(),name),nil

	
}

func HelloForMultiplePeople(names []string)(map[string]string ,error){
	messages:= make(map[string]string)

	for _,name:= range names{
		message,err:=RandomGreeting(name)
		if err !=nil{
			return nil,err
		}
		messages[name] = message
	}

	return messages,nil
}