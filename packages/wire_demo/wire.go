//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/google/wire"
)

// 初始化应用
func InitGreeter() *Greeter {
	wire.Build(NewUser, NewGreeter)
	return nil
}
