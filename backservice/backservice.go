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
	go p.Run(p.Session)
	return nil
}

func (p *Program) NetStatus() error {
	res, err := p.Session.GetInfo()
	if err != nil {
		p.Logger.Error(err)
	}
	if res.UserId != "" {
		p.Logger.Info(fmt.Sprintf("current user %s already login", res.UserId))
	} else {
		p.Logger.Info("logout.")
	}
	return nil
}

func (p *Program) RunOnce(session *fzuconnect.LoginSession) {
	// 检查当前用户是否登陆
	res, err := session.GetInfo()
	if err != nil {
		p.Logger.Error(err)
	}
	if res.UserId != session.Username {
		// 强制下线当前用户
		if res.UserId != "" {
			p.Logger.Info(fmt.Sprintf("logout %s", res.UserId))
			res, err := session.Logout()
			if err != nil {
				p.Logger.Error(err)
			}
			p.Logger.Info(fmt.Sprintf("logout result: %s", res.Result))
		}
		// 登陆目标用户
		{
			p.Logger.Info(fmt.Sprintf("%s login", session.Username))
			res, err := session.Connect()
			if err != nil {
				p.Logger.Error(err)
			}
			p.Logger.Info(fmt.Sprintf("result: %s", res.Result))
		}
	} else {
		p.Logger.Info(fmt.Sprintf("current user %s already login", res.UserId))
	}
}

// run 程序运行
func (p *Program) Run(session *fzuconnect.LoginSession) {
	p.RunOnce(session)
	for range time.Tick(time.Minute) { // 每隔1分钟登陆校园网
		p.RunOnce(session)
	}
}

// Stop 程序结束
func (p *Program) Stop(s service.Service) error {
	p.Logger.Info("stop service")
	return nil
}
