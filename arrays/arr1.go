package main

func Max(arr [5]int) (i,v int){
	maxi := -1
	maxval := -1
	for i,v := range arr{
		if v>maxval{
			maxi,maxval = i,v
		}
	}
	return maxi,maxval
}
func sum1 (arr [5]int) int{
	sumNums := 0
	for _,v := range arr{
		sumNums += v
	}
	return sumNums
}
func main(){
	arr := [5]int{21,22,13,41,32}
	i,v := Max(arr)
	println(i,v)

	sumNums := sum1(arr)
	println(sumNums)
}
