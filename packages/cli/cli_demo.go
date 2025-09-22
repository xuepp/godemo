package cli

import (
	"fmt"
	"log"
	"os"

	//文档地址:https://cli.urfave.org/v2/getting-started
	"github.com/urfave/cli/v2"
)

func Greet() {
	app := &cli.App{
		Name:  "greet",
		Usage: "fight the loneliness!",
		//Action 是当 CLI 程序运行时默认执行的函数
		Action: func(cCtx *cli.Context) error {
			fmt.Println("Hello friend!")
			//获取参数
			fmt.Println("Hello", cCtx.Args().Get(0), cCtx.Args().Get(1))
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Flags() {
	app := &cli.App{
		//定义了一个 全局标志（flag） --lang：
		//类型是字符串 (StringFlag)
		//默认值是 "english"
		//用途描述 "language for the greeting"
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "lang",
				Value: "english",
				Usage: "language for the greeting",
			},
		},
		Action: func(cCtx *cli.Context) error {
			name := "Nefertiti"
			if cCtx.NArg() > 0 {
				name = cCtx.Args().Get(0)
			}
			if cCtx.String("lang") == "spanish" {
				fmt.Println("Hola", name)
			} else {
				fmt.Println("Hello", name)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func Demo() {
	app := cli.NewApp()
	app.Name = "server"
	app.Version = "v1.0"
	app.Usage = "A test API service based on golang."
	app.Commands = []*cli.Command{}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
