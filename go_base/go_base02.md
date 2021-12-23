## 数组
   **数组的长度是固定的，所以在Go语言中很少直接使用数组**
 1. base
   * 定义: 同一种数据类型的固定长度的序列
   * 声明：var arrayname [lenth] type  eg: var arr1[5] int
   * for 循环遍历：
    ```
    for i:=0; i<len(arr1); i++{

    }
    for index, v:=index a{

    }
    ```
    * 数组是值类型，改变副本不会改变本身
 2. 一维数组
   ``` 
   d := [...]struct {
        name string
        age  uint8
    }{
        {"user1", 10}, // 可省略元素类型。
        {"user2", 20}, // 别忘了最后一行的逗号。
    }
    ```


## 变量
1. 定义变量
   var name type        若该变量未被赋值，go自动将其初始化，复制该变量类型的零值
2. 定义变量并赋值
   var name type = value    
3. 类型推断
   var name = value //go能够推断具有初始值的变量的类型，如果变量有初始值，那么可以省略type
4. 多值定义
   var name1， name2 type = value1， value2
   var name1， name2  = value1， value2
   var (
       name1 = value1
       name2 = value2
   )
5. 简短声明
   name := value   //要求name没有赋值，且左侧至少有一个变量尚未声明
6. go是强类型语言，不允许某一类行的变量赋值为其他类型的值


# 数据类型
1. bool
   true, false
2. 有符号整型
   int8
   int16
   int32
   int64
3. 无符号整型
   uint8
   uint16
   uint32
   uint64
4. 浮点型
   float32
   float64
5. 复数类型
   complex64    //实部和虚部都是float32类型的复数
   complex128   //实部和虚部都是float64类型的复数
   内建函数complex用于创建一个包含实部和虚部的复数
6. 其他数字类型
   byte(unit8)  //byte是unit8的别名
   rune(int32)  // rune是int32的别名
7. string类型
   name = “string”  //字节的集合
8.  类型转换
   type(value) // int(value)  将value转换成整型



# 切片
   * 定义：是对数组的一个连续片段的引用，所以切片是一个引用类型，这个片段可以是整个数组，也可以是由起始和终止索引标识的一些项的子集，需要注意的是，终止索引标识的项不包括在切片内。
   * slice [开始位置 : 结束位置]
         两者同时为 0 时，等效于空切片，一般用于切片复位
   * 从数组或切片生成新的切片
    ```
      var a = [3]int{1,2,3}
      fmt.Println((a,a[1:2]))
    ```
   * 直接声明新的切片
     每一种类型都可以拥有其切片类型，表示多个相同类型元素的连续集合，因此切片类型也可以被声明，切片类型声明格式如下：
     var name []Type
    ```
      var slice_str [] string
      var slice_int [] int
      var slice_int = [] int{}
    ```
   * make()生成切片
    如果需要动态地创建一个切片，可以使用 make() 内建函数，格式如下：
      make( []Type, size, cap ) size 指的是为这个类型分配多少个元素，cap 为预分配的元素数量，这个值设定后不影响 size，只是能提前分配空间，降低多次分配空间造成的性能问题。