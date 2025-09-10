##项目结构
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

