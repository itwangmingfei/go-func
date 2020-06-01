package main

import (
	"fmt"
	"log"
)


func main(){
	sli := []int{1,5,2,634,12,54,48}
	fmt.Println(sli)
	printSlice(firstDel(sli))
	printSlice(numDel(2,sli))
	printSlice(lastDel(sli))

}
//去掉第一个
func firstDel(sli []int) (int,[]int){
	value := sli[0]
	log.Printf("Del first slice value : %d \n",value)
	sli = sli[1:]
	return value,sli
}
//排除一个数据
func numDel(out int,sli []int)(int, []int) {
	if out > len(sli)-1{
		panic("nums must > len(sli)")
	}
	value := sli[out]
	log.Printf("Del %d slice value:%d \n",out,value)
	sli = append( sli[:out],sli[out+1:]...)
	return value,sli
}

func lastDel(sli []int) (int, []int){
	value := sli[len(sli)-1]
	log.Printf("Del last slice value:%d \n",value)
	sli = sli[:len(sli)-1]
	return value,sli
}

func printSlice(value int,slice []int){
	fmt.Println(value,slice)
}
