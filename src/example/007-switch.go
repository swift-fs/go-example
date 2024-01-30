package example

import (
	"fmt"
	"time"
)

func Switch() {
	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	now := time.Now()
	fmt.Println("now:", now)
	switch now.Weekday() {

	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	case time.Friday:
		fmt.Println("It's Friday")
	default:
		fmt.Println("It's a weekday")
	}

	// 不带表达式的 switch 是表达 if/else 逻辑的另一种方式。这里我们还展示了 case 表达式如何可以是非常量。
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() > 12:
		fmt.Println("It's after noon")
	}

	// 类型 switch 比较类型而不是值。您可以使用它来发现接口值的类型。在此示例中，变量 v 将具有与其子句相对应的类型

	whatType := func(i any) {
		switch v := i.(type) {
		case bool:
			fmt.Println("I'm a bool: ", v)
		case int:
			fmt.Println("I'm an int: ", v)
		default:
			fmt.Printf("Don't know type %T!\n", v)
		}
	}

	whatType(true)
	whatType(100)
	whatType(123.18)
	whatType("你好")

}
