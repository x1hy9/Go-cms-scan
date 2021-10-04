package main

import (
	"os"
	"scan/fileutil"
	"scan/json_new"
	"scan/network"

	cli "github.com/jawher/mow.cli"
	"github.com/remeh/sizedwaitgroup"
)

var (
	app     *cli.Cli
	threads = 10
	swg     sizedwaitgroup.SizedWaitGroup
)

func main() {
	//设置程序名称
	app = cli.App("扫描器", "欢迎使用扫描器")
	//-v 出版本号
	app.Version("v version", " 版本号1.0.0")

	//设置参数
	var (
		file    = app.StringOpt("f file", "", "文件名称")
		url     = app.StringOpt("u url", "", "单个URL")
		threads = app.IntOpt("t threads", 5, "线程数")
	)
	//设置spec
	app.Spec = "(-f | -u)"

	//执行的函数
	app.Action = func() {

		targetsSlice := make([]string, 0)
		if *threads > len(targetsSlice) {
			*threads = len(targetsSlice)
		}
		swg = sizedwaitgroup.New(*threads)
		//直接扫描
		if len(*url) != 0 {
			req, _ := network.Reqdata(*url)
			json_new.Detect(req)

		} else {
			filename, _ := fileutil.ReadFile(*file)
			targetsSlice = filename

			for i := 0; i < len(filename); i++ {
				swg.Add()
				// 开启一个并发
				go func(url string) {
					// 使用defer, 表示函数完成时将等待组值减1
					defer swg.Done()

					req, _ := network.Reqdata(url)
					json_new.Detect(req)

				}(filename[i])

			}

		}

		swg.Wait()

	}

	app.Run(os.Args)
}
