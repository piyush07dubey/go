package main

import "fmt"

func main(){
	var nums[4]int
	nums[0]=1
	fmt.Println(len(nums))
	fmt.Println(nums)

	var vals[4]bool
	vals[2]=true
	fmt.Println(vals)
	value:=[3]int{1,2,3}
	fmt.Println(value)

	new:=[2][2]int{{1,2},{3,4}}
	fmt.Println(new)

	
}