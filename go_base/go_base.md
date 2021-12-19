## 数组
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