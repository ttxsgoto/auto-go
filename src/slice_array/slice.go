package main

import "fmt"

func main() {
	var slice1 []int = make([]int, 10, 20)
	//slice1 := []string{ "123", "45", "asd"}   // 初始化索引的值，特定序列的值
	
	for i := 0; i < len(slice1); i++ {
		slice1[i] = 5 * i
	}
	fmt.Println(slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
	
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("s[:]", s[:])
	fmt.Println("s[2:7]", s[2:7], len(s[2:7]), cap(s[2:7]))
	// 新的slice对应的 len 为 7-2 即 j-i, cap为 10-2 即 原cap-i
	fmt.Println("s[1:]", s[1:], len(s[1:]), cap(s[1:])) // 9 9
	fmt.Println("s[:8]", s[:8], len(s[:8]), cap(s[:8])) // 8 10
	
	// slice[i:j:k] 来创建新的切片说明 , len=j-i, cap=k-i, k不能超过原始slice的cap大小
	new_s := s[2:6:10]
	new_s[2] = 22   // 原切片和新切片都是同一个底层数组，当改变任意一个数组的值，两个数组都改变
	fmt.Println("s[2:6:10]", new_s, len(new_s), cap(new_s))
	fmt.Println("s[:]", s[:])
	
	// append() 追加值到slice中, 容量可用存下新追加的元素
	new_s = append(new_s, 23, 24)
	//new_s = append(new_s, 24)
	fmt.Println("new_s append:", new_s) // [2 3 22 5 23 24]
	// 将原slice对应位置的值覆盖掉
	fmt.Println("s :", s)   // s : [0 1 2 3 22 5 23 24 8 9]
	
	// 当切片可用容量存不下新追加的元素时, 此时 s 和 new_ss 底层指向不同的数组
	// 切片的底层数组没有足够的可用容量， append()函数会创建一个新的底层数组，
	// 将原数组的值复制到新数组里，再追加新的值，不会影响原来的底层数
	s1 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	
	new_ss := s1[2:4]
	fmt.Println("new_ss", new_ss, len(new_ss), cap(new_ss))   // [2, 3] 2 8
	new_ss = append(new_ss, 33, 44, 55, 66, 77, 88, 99, 100)
	fmt.Println("append 9 num", new_ss, len(new_ss), cap(new_ss))   // [2 3 33 44 55 66 77 88 99 100 110] 11 16
	fmt.Println("s1 ", s1, len(s1), cap(s1)) // [0 1 2 3 4 5 6 7 8 9] 10 10
	
	// slice append slice
	s11 := []int{11, 22}
	s1 = append(s1, s11...)
	fmt.Println("s1", s1, len(s1))
	
	for i, v := range s1 {
		fmt.Println(i, v)
	}
	
	var nils []int
	
	// 空切片,已初始化
	s2 := []int{}
	s3 := make([] int, 0)
	fmt.Println("s2", s2, len(s2), cap(s2))
	fmt.Println("s3", s3, len(s3), cap(s3))
	if nils == nil {
		fmt.Println("nils", nils, len(nils), cap(nils))
	}
	s4 := make([]int, 5)
	s4 = append(s4, 1, 2, 3)
	fmt.Println(s4)
	s5 := make([]int,0)
	s5 = append(s5,1,2,3,4)
	fmt.Println(s5)
}
