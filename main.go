// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	hlog.SetLevel(hlog.LevelDebug)
	app := cli.NewApp()
	app.Name = "test_quck"
	app.Usage = "测试"
	app.Version = "1.0"

	// Commands
	app.Commands = []*cli.Command{
		{
			Name:  "web",
			Usage: "test ",
			Action: func(context *cli.Context) error {
				f, err := os.OpenFile("./output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
				if err != nil {
					panic(err)
				}
				defer f.Close()
				hlog.SetOutput(f)
				h := server.Default(
				//server.WithHostPorts(config.Config.Port),
				)
				register(h)
				h.Spin()
				return nil
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		hlog.Info(err)
	}

}
