package main

import "fmt"
import "time"

func main(){
	i:=5
	switch i{
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	default:
		fmt.Println("other")
	}
	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("working day")
	}
	whoAmi:=func (i interface{})  {
		switch j:=i.(type){
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string",j)
		case float64:
			fmt.Println("float64",j)
		default:
			fmt.Println("other",j)
		}
	}
	whoAmi(1)
	whoAmi("golang")
	whoAmi(1.2345)
}