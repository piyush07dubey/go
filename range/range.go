package main

import "fmt"

func main(){
	num:=[]int{2,4,6,8,10}
	for i,num:=range num{
		fmt.Println(num,i)
	}
	m:=map[string]string{"fname":"john","lname":"doe"}
	for k,v:=range m{
		fmt.Println(k,v)
	}
}