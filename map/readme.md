#### 定义
- map 是一种特殊的数据结构：一种元素对（pair）的无序集合，pair 的一个元素是 key，对应的另一个元素是 value，所以这个结构称为字典
- map为引用类型，内存用 make 方法来分配
- map长度是可以动态增长，未初始化的 map 的值是 nil
- 不要使用 new，永远用 make 来构造 map


#### 声明
    var map1 map[keytype]valuetype
    var map1 map[string]int

map 的初始化：var map1 = make(map[keytype]valuetype)


#### map容量
- 标明 map 的初始容量 capacity
- make(map[keytype]valuetype, cap) map2 := make(map[string]float32, 100)








