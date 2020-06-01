package main

import "fmt"


func main(){
	var maps map[string]int
	maps =make(map[string]int)
	maps["name"] = 12
	maps["name1"] = 1
	maps["name2"] = 2
	maps["name3"] = 3
	maps["name4"] = 4
	fmt.Println(maps)
	/*遍历无序排列*/
	//删除一个key
	delete(maps,"name4")
	for k,v := range maps{
		fmt.Println(k,v)
	}
	name3,ok :=maps["name23"]
	fmt.Println(name3,ok)
}

