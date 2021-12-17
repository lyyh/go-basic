package modify

import "fmt"

func Modify() {
	m := make(map[string]int)
	m["Anser"] = 100
	fmt.Println(m["Anser"])
	m["Anser"] = 1
	fmt.Println(m["Anser"])
	delete(m, "Anser")
	v, ok := m["Anser"]
	fmt.Println(v, ok)
}
