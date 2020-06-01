package main

import "fmt"

//切片获取最大值
func MaxValue(arr []int) (i, v int) {
	maxi := 0
	maxValue := -1
	for i, val := range arr {

		if maxValue < val {
			maxi, maxValue = i, val
		}
	}
	return maxi, maxValue
}

func sum (arr []int) int{
	sumNums := 0
	for _,v := range arr{
		sumNums += v
	}
	return sumNums
}
//slice  视图会修改原模型数值
func setNums(arr []int) {
	arr[0] = 100
}

func main() {
	arr := []int{1, 2, 4, 3, 5, 9, 7}
	i, v := MaxValue(arr)
	fmt.Println(i, v)

	sumNums := sum(arr)
	println(sumNums)
	setNums(arr)
	fmt.Println(arr[2:5])
}
