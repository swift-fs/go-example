package example

import (
	"fmt"
	"maps"
)

func Map() {

	// 当使用 fmt.Println 打印时，地图以 map[k:v k:v] 形式显示。
	// To create an empty map, use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int)

	// Set key/value pairs using typical name[key] = val syntax.
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// Get a value for a key with name[key].
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// If the key doesn’t exist, the zero value of the value type is returned.
	v3 := m["k3"]
	fmt.Println("v3:", v3)

	// The builtin len returns the number of key/value pairs when called on a map.
	fmt.Println("len:", len(m))

	// The builtin delete removes key/value pairs from a map.
	delete(m, "k2")
	fmt.Println("map:", m)

	// To remove all key/value pairs from a map, use the clear builtin.
	clear(m)
	fmt.Println("map:", m)

	// 从映射获取值时可选的第二个返回值指示该键是否存在于映射中。这可用于消除缺少的键和具有零值的键（例如 0 或 "" ）之间的歧义。这里我们不需要值本身，所以我们用空白标识符 _ 忽略它。
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 定义和初始化一起
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	// maps 包包含许多有用的实用函数。
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("n == n2")
	}
}
