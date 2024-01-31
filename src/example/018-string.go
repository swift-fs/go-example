package example

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	// Go语言字符串是一个只读的字节片
	// 在Go语言中，字符的概念被称为 rune -它是一个整数，代表一个Unicode码位
	// 标准库对字符串进行了特殊处理--将其作为UTF-8编码的文本的容器
	// 字符串等价于 []byte ，这将产生存储在其中的原始字节的长度。
	const s = "你好呀"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	// 要计算字符串中有多少个符文，我们可以使用 utf8 包
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// range 循环专门处理字符串，并将每个 rune 沿着及其在字符串中的偏移量进行解码。
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	// 我们可以通过显式地使用 utf8.DecodeRuneInString 函数来实现相同的迭代。
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		// 用单引号括起来的值是rune字面值。我们可以直接将 rune 值与rune字面量进行比较。
		fmt.Println("runeValue == 你", runeValue == '你')
	}
}
