package hello

import (
	"fmt"
)

// pass By value
// Pass by Reference
// Pointers basic concepts
var value string = "Lemon"

func Hello() {
	ptr := &value
	
	fmt.Println("The value address is:",ptr) //reference
	fmt.Println("The actual value:",*ptr) //dereference

	fmt.Println("value before caling the function", value)
	passByValue(value)
	fmt.Println("value after caling the function", value)
	fmt.Println("pass by reference ...........................")
	passByReference(&value)
	fmt.Println("address",&value)
	fmt.Println("value after caling the function and pass by reference", value)

}

func passByValue(_value string) {
	_value = "Yollow"
	fmt.Println("inside the function",_value)
}

func passByReference(param *string){
	*param = "yesNewName"
	
}