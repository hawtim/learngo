package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	// 变量定义
	// x := 5
	// y := 7
	// sum := x + y
	// fmt.Println("hello, world!", sum)

	// if else 条件语句
	// if sum >= 12 {
	// 	fmt.Println("More than to 12")
	// } else if x <= 5 {
	// 	fmt.Println("x less than to 5")
	// }

	// 可拓展数组
	// a := [] int {1, 2, 3, 4, 5}
	// a = append(a, 13)
	// fmt.Println("hello, world!", a)

	// 可拓展 map
	// vertices := make(map[string]int)
	// vertices["triangle"] = 2
	// vertices["square"] = 3
	// vertices["dodecagon"] = 12
	// delete(vertices, "triangle")
	// fmt.Println(vertices)

	// for 循环
	// for i:= 0; i < 5; i++ {
	// 	fmt.Println(i)
	// }

	// 遍历数组
	// arr := []string{"a", "b", "c"}
	// for index, value := range arr {
	// 	fmt.Println(index, value)
	// }

	// 遍历 map
	// m := make(map[string]string)
	// m["a"] = "alpha"
	// m["b"] = "beta"
	// for key, value := range m {
	// 	fmt.Println(key, value)
	// }

	// 函数调用
	// result := sumfunc(1, 2)
	// fmt.Println(result)

	// 回调函数
	// result, err := sqrt(-16)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(result)
	// }
	
	// 构造函数/构造类

	p := person{name: "hawtim", age: 27}
	fmt.Println(p.age)

	// 内存地址
	i := 7
	inc(&i)
	fmt.Println(i, &i)
}

func sumfunc(x int, y int) int {
	return x + y
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}
	return math.Sqrt(x), nil
}

type person struct {
	name string
	age int
}

func inc(x *int) {
	*x++
}