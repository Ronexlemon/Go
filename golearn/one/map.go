package main
import (
	"fmt"
)

//sytax
/*
make(map[key-tpe]value-type)
*/
func mapBasic()map[string]int{
m:=make(map[string]int)
m["John"]=20
m["doe"] =25
m["kim"] =30
m["jane"] =35
fmt.Println("map:",m)
return m

}

func getAvalueFromMap(){
	m:=mapBasic()
	v1 :=m["john"]
	fmt.Println("John:",v1)
	fmt.Println("length of the map",len(m))//return the number of key/value pairs

}

func returnZeroIfKeyDoesNotExists(){
	m:= mapBasic()
	v1 := m["ron"]
	fmt.Println("value does not exists",v1)
}

//remove key/value pairs from a map

func deleteKeyValuePairs(){
	m:=mapBasic()
	delete(m,"kim")
	fmt.Println("delete pairs with key kim:",m)
}

//to remove all keys pairs from a map
func clearAllKeysValues(){
	m:=mapBasic()
	clear(m)
	fmt.Println("the mapis empty",m)
}

//The optional second return value when getting a value from a map indicates if the key was present in the map. This can be used to disambiguate between missing keys and keys with zero values like 0 or "". 
func noNeedValueItself(){
	m:=mapBasic()
	_,prs:=m["janes"]
	fmt.Println("prs: ",prs)//returns true if the key was present and false if not
}

//declare and initialize new map
func initNewMapOnSameLine(){
	m:=map[string]int{"homa":5,"johndoe":5}
	fmt.Println("before adding to map",m)
	m["hun"]=7
	fmt.Println("after adding to map:",m)
}
func Map(){
	getAvalueFromMap()
	returnZeroIfKeyDoesNotExists()
	deleteKeyValuePairs()
	clearAllKeysValues()
	noNeedValueItself()
	initNewMapOnSameLine()
	
}