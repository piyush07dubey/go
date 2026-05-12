package main

import (
	"fmt"
	"maps"
)

func main(){
	m:=make(map[string]string)
	m["name"]="golang"
	m["version"]="1.23"
	fmt.Println(m)
	fmt.Println(m["name"])

	n:=make(map[string]int)
	n["key1"]=10
	fmt.Println(n)

	f:=map[string]int{"price":40,"phone":3}
	fmt.Println(f)
	k,ok:=f["price"]
	if ok {
		fmt.Println("all ok",k)
	}else{
		fmt.Println("not ok")
	}
	x:=map[string]int{"price":40,"phone":3}
	y:=map[string]int{"price":40,"phone":3}
	fmt.Println(maps.Equal(x,y))
	
}