# GoDemo

这是一个用于学习和演示 Go 语言核心概念的项目。  
包括函数、方法、数据类型、结构体、接口、并发、错误处理、包管理等模块。

## 项目结构
```
go-concepts-demo/
├─ go.mod                     # Go 模块文件
├─ main.go                    # 程序入口
├─ README.md                  # 项目说明
├─ pkg/                       # 核心概念模块（可复用）
│   ├─ functions/             # 函数和方法
│   │   ├─ function_demo.go
│   │   └─ method_demo.go
│   ├─ types/                 # 数据类型、结构体、接口
│   │   ├─ struct_demo.go
│   │   ├─ interface_demo.go
│   │   └─ type_demo.go
│   ├─ concurrency/           # 并发相关
│   │   ├─ goroutine_demo.go
│   │   ├─ channel_demo.go
│   │   └─ select_demo.go
│   ├─ packages/              # 包与导入
│   │   └─ package_demo.go
│   ├─ errors/                # 错误处理
│   │   └─ error_demo.go
│   ├─ generics/              # 泛型
│   │   └─ generics_demo.go
│   ├─ io/                    # 文件操作、输入输出
│   │   └─ io_demo.go
│   └─ others/                # 其他概念（指针、defer、recover 等）
│       └─ misc_demo.go
├─ examples/                  # 可运行示例脚本
│   ├─ basic_usage.go
│   └─ advanced_usage.go
└─ docs/                      # 文字说明或学习笔记

```

## 模块设计说明
### pkg/functions/

- function_demo.go：函数声明、匿名函数、闭包、可变参数
- method_demo.go：方法、值接收者 vs 指针接收者

### pkg/types/
- type_demo.go：基本类型（int, string, bool）、常量、枚举
- struct_demo.go：结构体、嵌套结构体、构造函数
- interface_demo.go：接口定义、接口实现、多态演示

### pkg/concurrency/
- goroutine_demo.go：启动 goroutine、goroutine 调度
- channel_demo.go：channel 的创建、发送、接收
- select_demo.go：select 多路复用、超时案例

### pkg/packages/
- package_demo.go：自定义包、标准库包、导出规则（首字母大小写）
- pkg/errors/
- error_demo.go：内置错误、自定义错误、panic/recover

### pkg/generics/
- generics_demo.go：Go 1.18+ 泛型使用示例

### pkg/io/
- io_demo.go：文件读写、标准输入输出

### pkg/others/
- misc_demo.go：指针、defer、recover、类型转换、反射等零散知识点

### examples/
- basic_usage.go：组合调用各个模块，演示基本用法
- advanced_usage.go：组合复杂场景，例如并发 + 错误处理 + 泛型