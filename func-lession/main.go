package main // 声明 main 包，表明当前是一个可执行程序

// 导入内置 fmt 包

// func intSum(x, y int) int {
// 	return x + y
// }

// // 可变参数
// func intSum2(x int, args ...int) (int, int) {
// 	var sum = x
// 	for _, v := range args {
// 		sum = sum + v
// 	}
// 	return x, sum
// }

// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil
// 	}
// 	// return x
// }

// type calculation func(int, int) int

// var c calculation

// 高阶函数
func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	// sum := intSum(1, 2)
	// fmt.Sprintln(sum)
	// _, total := intSum2(1, 2, 3, 4, 5)
	// fmt.Sprint(total)
	// res1 := calc(1, 2, add)
	// fmt.Println(res1)
	// add: = func (x,y init)  {

	// }
}
