package main

import (
	"flag"
	"fzuconnect/backservice"
	"log"
	"os"
	"path/filepath"

	"github.com/kardianos/service"
)

// Config 配置文件地址
type Config struct {
	config *string
}

var path, _ = filepath.Abs(filepath.Dir(os.Args[0]))

var config *Config = &Config{
	config: flag.String("c", "config.ini", "config path. default: config.ini"), // 服务必须是绝对路径
}

func main() {
	flag.Parse()
	if !filepath.IsAbs(*config.config) { // 如果不是绝对路径，转换为绝对路径
		*config.config = filepath.Join(path, *config.config)
	}
	// 注册服务
	program := new(backservice.Program)
	program.LoadIni(*config.config)
	var args []string
	if len(flag.Args()) > 1 {
		args = flag.Args()[1:]
	}
	svcConfig := &service.Config{
		Name:        program.Name,
		DisplayName: program.DisplayName,
		Description: program.Description,
		Arguments:   args, // 注册服务时去掉第一个参数
	}
	s, err := service.New(program, svcConfig)
	if err != nil {
		log.Fatal(err)
	}
	program.Logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err != nil {
			program.Logger.Error(err)
		}
	}()
	if len(flag.Args()) > 0 {
		switch flag.Args()[0] {
		case "install":
			err = s.Install()
			return
		case "uninstall":
			err = s.Uninstall()
			return
		case "stop":
			err = s.Stop()
			return
		case "restart":
			err = s.Restart()
			return
		case "start":
			err = s.Start()
			return
		case "status":
			err = program.NetStatus()
			return
		}
	}
	err = s.Run()
}
