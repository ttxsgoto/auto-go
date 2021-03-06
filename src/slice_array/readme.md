#### 数组说明
- 数组是具有相同 唯一类型 的一组已编号且长度固定的数据项序列
- 数组长度必须是一个常量表达式，并且必须是一个非负整数
- 数组元素可以通过 索引（位置）来读取（或者修改）

https://github.com/unknwon/the-way-to-go_ZH_CN/blob/master/eBook/07.1.md


#### 将数值传递给函数
把一个大数组传递给函数会消耗很多内存,有两种方法可以避免(array_sum)：
- 传递数组的指针
- 使用数组的切片

#### 切片说明
- 切片（slice）是对数组一个连续片段的引用
- 切片是一个引用类型，类似于python中的列表
- 切片是一个 长度可变的数组
- 切片是可索引的，并且可以由 len() 函数获取长度
- 和数组不同的,切片的长度可以在运行时修改
- 不要用指针指向 slice。切片本身已经是一个引用类型，它本身就是一个指针

切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率，所以在 Go 代码中 切片比数组更常用

声明
var identifier []type ,不需要指定长度，未初始化之前默认为nil，长度为0

##### 使用make()创建切片
当相关数组还没有定义时，我们可以使用 make() 函数来创建一个切片同时创建好相关数组
var slice1 []type = make([]type, len)
slice1 := make([]type, len)
make 接受 2 个参数：元素的类型以及切片的元素个数
如果你想创建一个 slice1，它不占用整个数组，而只是
占用以 len 为个数个项，那么只要：slice1 := make([]type, len, cap)
make 的使用方式是：func make([]T, len, cap)，其中 cap 是可选参数

##### new() 和 make() 的区别

- new(T) 为每个新的类型T分配一片内存，初始化为 0 并且返回类型为*T的内存地址：这
  种方法 返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体
- make(T) 返回一个类型为 T 的初始值，它只适用于3种内建的引用类型：切片、map 
  和 channel
- new 函数分配内存，make 函数初始化
    var v []int = make([]int, 10, 50)
    v := make([]int, 10, 50)

    
##### 切片复制和追加
- func append(s[]T, x ...T) []T 其中 append 方法将 0 个或多个具有相同类型 s 的元素追加到切片后面并且返回新的切片
- 追加的元素必须和原切片的元素同类型
- s 的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储


























