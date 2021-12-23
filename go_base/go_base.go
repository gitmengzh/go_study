package main //每个 Go 源文件必须先声明它所属的包

import ( //导入程序中所依赖的包
	"fmt"
)

func main() {

	hellowrod()
	variable()
	arr1()
	muil_arr()
	slice1()
}

func hellowrod() { // first go
	fmt.Println("Hello world!")
}

func variable() { //变量使用

	var age int //变量声明
	fmt.Println("my age is", age)
	age = 10 //赋值
	fmt.Println("Tom age is", age)
	var age2 int = 30 // 声明变量并赋值
	fmt.Println("Lily age is", age2)
	var age3 = 25 // 推断类型
	fmt.Println("tomy age is", age3)
	var name, age4 = "TOO", 18 // 多变量声明赋值
	fmt.Println(name, "age is", age4)
	name1 := "SONY"
	fmt.Println(name1, "TV")

}

func arr1() {
	var arr1 [5]int
	for i := 1; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	fmt.Println(arr1) //[0 2 4 6 8]
	arr3 := [...]int{1, 2, 3}

	for k, v := range arr3 { //遍历数组
		fmt.Println(k, v)
	}

}

func muil_arr() {
	var arr0 [5][3]int = [...][3]int{{}, {}, {}, {1, 2, 3}, {4, 5, 6}}
	var arr2 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}
	fmt.Println(arr2, arr0)
	// 声明并初始化数组中索引为 1 和 3 的元素
	var arr4 = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	arr5 := [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(arr4, arr5)
}

func slice1() {
	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:])

}
