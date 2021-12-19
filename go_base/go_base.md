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