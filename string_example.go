package main

import (
	"fmt"
	"strings"
	"strconv"
)

// 字符索引、字符是否存在，字符统计， 字符大小写， 字符删除前后空格， 分割字符串， 拼接字符串

func main() {
	var str string = "Hi, I'm ttxsgoto, Hi."
	var manyG = "gggggggggg"

	var origS string = "Hi there! "
	var newS string

	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The position of \"ttxsgoto\" is:")
	fmt.Printf("%d\n", strings.Index(str, "ttxsgoto"))

	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))

	fmt.Printf("Number of H's in %s is: ", str)
	fmt.Printf("%d\n", strings.Count(str, "H"))

	fmt.Printf("Number of double g's in %s is: ", manyG)
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))

	newS = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)

	splitjoin()
	conversion()
}

func splitjoin() {
	str := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(str)
	fmt.Printf("Splitted in slice: %v\n", sl)
	fmt.Println()

	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for index, val := range sl2 {
		fmt.Printf("%s - %d\n", val, index)
	}
	fmt.Println()
	str3 := strings.Join(sl2, ";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)
}

func conversion() {

	var orig string = "666"
	var an int
	var newS string
	fmt.Println()
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)        // 操作系统平台下 int 类型所占的位数
	an, _ = strconv.Atoi(orig)	// 将字符串转换为 int 型
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)	// 数字 i 所表示的字符串类型的十进制数
	fmt.Printf("The new string is: %s\n", newS)

}







