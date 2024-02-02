package multimodule

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func Hello_Multimodule() {
	fmt.Println(reverse.String("Yollow"))

}
