package wire

import "fmt"

// User 是一个简单的业务对象
type User struct {
	Name string
}

// NewUser 是 User 的构造函数
func NewUser() *User {
	return &User{Name: "Tom"}
}

// Greeter 依赖 User
type Greeter struct {
	User *User
}

// NewGreeter 是 Greeter 的构造函数
func NewGreeter(u *User) *Greeter {
	return &Greeter{User: u}
}

// Greet 打印问候
func (g *Greeter) Greet() {
	fmt.Printf("Hello, %s!\n", g.User.Name)
}
