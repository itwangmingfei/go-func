package main

import "fmt"

func main(){
	s1 := []int{1,2,3,4,5}
	var s2 []int
	s2 = make([]int,5)
	copy(s2,s1)
	fmt.Println(s2)
}
