package main

import "go-example/src/example"

// go run 001-hello-world.go 开发期间编译运行
// go build -o hello 001-hello-world.go  生产构建产物,指定产物名字为hello
func main() {

	example.Hello()
	example.Values()
}
