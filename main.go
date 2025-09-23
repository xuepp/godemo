package main

import "github.com/xuepp/godemo/packages/cli"

//import "github.com/xuepp/godemo/packages/cli"

func main() {
	//functions.FunctionDemo()
	//functions.MethodDemo()
	//types.StructDemo()
	//types.ArrayDemo()
	//types.SliceDemo()
	//types.InterfaceDemo()
	//cli.Greet()
	//cli.Flags()
	cli.Demo()

}

// 定义一个打印器
// type Printer struct{}

// func (p *Printer) Print(msg string) {
// 	fmt.Println(msg)
// }

// // 消费者
// type Greeter struct{}

// func (g *Greeter) Greet() {
// 	p := &Printer{} // 硬编码依赖
// 	p.Print("Hello, World!")
// }

// func main() {
// 	g := &Greeter{}
// 	g.Greet()
// }
