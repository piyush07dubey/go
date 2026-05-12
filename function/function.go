package main

import "fmt"
func add(a int,b int)int{
	return a+b
}
func getlanguage()(string,string,string,int){
	return "golang","javascript","c",6
}




func main(){
result:=add(1,2)
fmt.Println(result)
fmt.Println(getlanguage())

}