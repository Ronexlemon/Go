package main


import ("fmt")
//go supports for anonymous functions which form closures

func  InSeq() func() int{
	i:=0
	return func() int{
		i++
		return i
	}
}

func GetName() func(name string)string{
	return func(name string) string {
		return name
	}

}

//anonymous define using a varibale
func defineWithVariable(){
	my:= func()string{
		return "ooh my anonoymous"

	}
	my()
}

func Closure(){
	inseq:= InSeq()
	fmt.Println(inseq())
	fmt.Println(inseq())

	getname:= GetName()

	fmt.Println(getname("yollow who"))
	defineWithVariable()
}