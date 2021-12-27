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
	slice2()
	slice3()
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
	fmt.Println(a, a[1:])            // 对数组a进行切片
	var slice_int = []int{1, 2, 3}   // 声明一个全新的切片，类型为int
	var slice_string []string        //声明一个全新的切片，类型为string
	slice_make := make([]int, 2, 10) //使用make（）函数声明一个切片
	fmt.Println("slice_int:", slice_int, "slice_string:", slice_string, "slice_make:", slice_make)

}

func slice2() { //append()
	a := make([]int, 2, 10)
	a = append(a, 1, 3, 5) // 追加多个元素, 手写解包方式
	fmt.Println(a)
	var b []int
	b = append(b, []int{2, 4, 6}...) // 追加一个切片, 切片需要解包
	fmt.Println(b)
	var c = []int{5, 6, 7}
	c = append([]int{4}, c...) //在开头添加一个元素
	fmt.Println("c添加一个元素：", c)
	c = append([]int{1, 2, 3}, c...) //在开头添加一个切片
	fmt.Println("c添加一个切片：", c)
	var d = []int{1, 2, 3, 6, 9}
	d = append(d[:3], append([]int{4}, d[3:]...)...) // 在下标3处添加一个元素
	fmt.Println("在第3个下标处，插入一个元素", d)
	d = append(d[:5], append([]int{7, 8}, d[5:]...)...) //在下标5处，添加一个切片
	fmt.Println("在第5个下标下，添加一个切片", d)

}

func slice3() { //copy()
	var srca = []int{1, 2, 3}
	var desta = []int{4, 5, 6, 7}
	copy(desta, srca)
	fmt.Println("拷贝后切片：", desta) // 拷贝源切片中的所有3个元素，替换掉目的切片的前3个元素
	var srcb = []int{1, 2}
	var destb []int
	copy(destb, srcb)
	fmt.Println(destb) //拷贝后结果为空，因为目的切片内没有元素
	srcc := []int{1, 2, 3}
	destc := make([]int, 2, 10)
	copy(destc, srcc)
	fmt.Print(destc) //根据目标切片内元素个数，而非提前分配长度进行拷贝
}
