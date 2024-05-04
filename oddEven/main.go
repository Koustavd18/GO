package main

import (
"fmt"
)

func main(){

var nums []int

for i:=0; i<=10; i++{
	nums = append(nums,i)
}

for _, n:= range nums{
	if n%2 == 0 {
fmt.Println(n," is Even")
} else {
fmt.Println(n," is Odd")
}
}
}
