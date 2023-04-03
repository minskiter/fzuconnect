package backservice

import (
	"fmt"
	fzuconnect "fzuconnect/fzulogin"
	"time"

	"github.com/kardianos/service"
	"gopkg.in/ini.v1"
)

// Program 后台服务
type Program struct {
	Name           string
	DisplayName    string
	Description    string
	ConfigFileName string
	Logger         service.Logger
	Session        *fzuconnect.LoginSession
}

// LoadIni 加载配置文件
func (p *Program) LoadIni(filename string) error {
	p.ConfigFileName = filename
	cfg, err := ini.Load(filename)
	if err != nil {
		return err
	}
	p.Name = cfg.Section("service").Key("name").String()
	p.DisplayName = cfg.Section("service").Key("displayname").String()
	p.Description = cfg.Section("service").Key("description").String()
	return nil
}

// Start 程序开始
func (p *Program) Start(s service.Service) error {
	p.Logger.Info(fmt.Sprintf("%v start", p.DisplayName))
	p.Session = new(fzuconnect.LoginSession)
	p.Session.LoadIni(p.ConfigFileName)
	go p.run(p.Session)
	return nil
}

// run 程序运行
func (p *Program) run(session *fzuconnect.LoginSession) {
	res := session.Connect()
	p.Logger.Info(res)
	for range time.Tick(time.Minute * 5) { // 每隔5分钟登陆校园网
		res := session.Connect()
		p.Logger.Info(res)
	}
}

// Stop 程序结束
func (p *Program) Stop(s service.Service) error {
	p.Logger.Info("结束服务")
	return nil
}
