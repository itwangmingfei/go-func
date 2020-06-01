package main

import (
	"fmt"
	"os"
)

//读取文件配置
func readMaze(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic("文件读取失败" + path)
	}
	defer file.Close()
	var row, col int
	//获取行列
	fmt.Fscanf(file, "%d %d\n", &row, &col)
	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			if j == col-1 {
				fmt.Fscanf(file, "%d\n", &maze[i][j])
			} else {
				fmt.Fscanf(file, "%d", &maze[i][j])
			}
		}
	}
	return maze
}

//设置点
type SetPoint struct {
	i, j int
}

//设置移动的点 上左下右

var goPoint = [4]SetPoint{
	{-1, 0}, {0, -1}, {0, 1}, {1, 0},
}

//判断有没有超出范围
func (p SetPoint) at(maze [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(maze) {
		return 0, false
	}
	if p.j < 0 || p.j >= len(maze[p.i]) {
		return 0, false
	}
	return maze[p.i][p.j], true
}

//移动
func (p SetPoint) add(pt SetPoint) SetPoint {
	return SetPoint{p.i + pt.i, p.j + pt.j}
}

func walk(maze [][]int, start, end SetPoint) [][]int {

	//定义一个和迷宫一样的锚点，都是0值
	setSteps := make([][]int, len(maze))
	for i := range setSteps {
		setSteps[i] = make([]int, len(maze[i]))
	}
	//获取到的点存入这里
	Q := []SetPoint{start}
	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]
		if cur == end {
			break
		}
		for _, onePoint := range goPoint {
			next := cur.add(onePoint)
			val, ok := next.at(maze)
			//是否有墙
			if !ok || val == 1 {
				continue
			}
			//去掉访问的点
			val, ok = next.at(setSteps)
			if !ok || val != 0 {
				continue
			}
			//去掉起点
			if next == start {
				continue
			}
			//设置setsteps 值
			val, _ = cur.at(setSteps)
			setSteps[next.i][next.j] = val + 1
			Q = append(Q, next)

		}
	}

	return setSteps
}

func main() {
	//1,读取maze配置文件
	path, _ := os.Getwd()
	maze := readMaze(path + "/maze/maze.ini")
	steps := walk(maze, SetPoint{0, 0}, SetPoint{len(maze) - 1, len(maze[0]) - 1})
	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
