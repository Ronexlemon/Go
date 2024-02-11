package hello

import (
	"fmt"
)

// pass By value
// Pass by Reference
// Pointers basic concepts
var value string = "Lemon"

func Hello() {
	fmt.Println("value before caling the function", value)
	passByValue(value)
	fmt.Println("value after caling the function", value)

}

func passByValue(_value string) {
	_value = "Yollow"
}
