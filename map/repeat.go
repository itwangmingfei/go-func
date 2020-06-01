package main

import (
	"fmt"
)

type list struct {
	i,start int
}
//获取不重复的数据和内容
func repeatSubstr(s string) (int,[]list) {
	setmap := make(map[rune]int)
	start := 0
	length := 0
	var newlist list
	var lists []list
	for i, value := range []rune(s) {
		//start位置定义
		lastk, ok := setmap[value]
		if ok && lastk >= start {
			start = lastk + 1
		}
		//length长度计算
		if i-start+1 > length {
			length = i - start + 1
		}
		//存入setmap
		setmap[value] = i
		newlist.i = i
		newlist.start = start
		lists= append(lists,newlist)
	}
	return length,lists
}
func main() {
	s := "我爱我你我他"
	length ,list := (repeatSubstr(s))
	fmt.Println("获取长度",length)
	bytes := []rune(s)
	for _,row := range list {
		if row.i-row.start+1 == length{
			for i := row.start;i< row.i+1 ;i++  {
				fmt.Print(string(bytes[i]))
			}
			fmt.Println()
		}
	}
}
