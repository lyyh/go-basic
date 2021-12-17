package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	Lat, Long float64
}

func declValue() {
	var b map[string]int
	b = make(map[string]int, 10)
	fmt.Printf("%#v", b)
}

// 按照指定顺序遍历map
func traverseMap() {

}

func create() {
	m := map[string]int{
		"username": 1,
		"age":      2,
	}
	fmt.Printf("%#v", m)

	delete(m, "username")
	if v, ok := m["张三"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("查无此人")
	}
}

func printSort() {
	var map1 = make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	// 无序输出
	for key, val := range map1 {
		fmt.Printf("%v %v\n", key, val)
	}
	var sli = []int{}
	for i := range map1 {
		sli = append(sli, i)
	}
	sort.Ints(sli)
	for i := 0; i < len(sli); i++ {
		fmt.Println(map1[sli[i]])
	}
}

func sortMap() {
	printSort()

	// rand.Seed(time.Now().UnixNano())
	// scopeMap := make(map[string]int, 100)
	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%d", i)
	// 	value := rand.Intn(100)
	// 	scopeMap[key] = value
	// }

	// keys := make([]string, 0, 200)
	// for key := range scopeMap {
	// 	keys = append(keys, key)
	// }

	// sort.Strings(keys)
	// for _, key := range keys {
	// 	fmt.Println(key, scopeMap[key])
	// }
	// rand.Seed(time.Now().UnixNano()) //初始化随机数种子

	// var scoreMap = make(map[string]int, 200)

	// for i := 0; i < 100; i++ {
	// 	key := fmt.Sprintf("stu%d", i) //生成stu开头的字符串
	// 	value := rand.Intn(100)        //生成0~99的随机整数
	// 	scoreMap[key] = value
	// }
	// //取出map中的所有key存入切片keys
	// var keys = make([]string, 0, 200)
	// for key := range scoreMap {
	// 	keys = append(keys, key)
	// }
	// //对切片进行排序
	// sort.Strings(keys)
	// //按照排序后的key遍历map
	// for _, key := range keys {
	// 	fmt.Println(key, scoreMap[key])
	// }
}
func main() {
	// modify.Modify()
	// declValue()
	// create()
	sortMap()
	// v := map[string]Vertex{
	// 	"Bill": {
	// 		0.1,
	// 		0.2,
	// 	},
	// }
	// fmt.Printf("%v", v)
}
