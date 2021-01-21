// Package fzuconnect 福州大学校园网连接工具
package fzuconnect

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/ini.v1"
)

// LoginSession 登陆结构
type LoginSession struct {
	Username string    // 用户名
	Password string    // 密码
	Config   *ini.File // 配置文件
}

// LoginResponse 响应结构体
type LoginResponse struct {
	UserIndex         string `json:"userIndex"`
	Result            string `json:"result"`
	Message           string `json:"message"`
	Forwordurl        string `json:"forwordurl"`
	KeepaliveInterval int    `json:"keepaliveInterval"`
	ValidCodeURL      string `json:"validCodeUrl"`
}

// LoadIni 加载配置文件
func (session *LoginSession) LoadIni(filename string) error {
	var err error
	session.Config, err = ini.Load(filename)
	if err != nil {
		return err
	}
	session.Password = session.Config.Section("common").Key("password").String()
	session.Username = session.Config.Section("common").Key("username").String()
	return nil
}

// ReloadIni 重载配置文件
func (session *LoginSession) ReloadIni() error {
	if session.Config == nil {
		panic("配置文件为空")
	}
	err := session.Config.Reload()
	if err != nil {
		return err
	}
	session.Password = session.Config.Section("common").Key("password").String()
	session.Username = session.Config.Section("common").Key("username").String()
	return nil
}

// Connect 连接到校园网
func (session *LoginSession) Connect() LoginResponse {
	client := &http.Client{
		Timeout: 2 * time.Second,
	}
	req, _ := http.NewRequest("POST", "http://210.34.48.49/eportal/InterFace.do", nil)
	// 设置传输的数据
	query := req.URL.Query()
	query.Add("method", "login")
	query.Add("userId", session.Username)
	query.Add("password", session.Password)
	query.Add("queryString", "wlanuserip=4168a23cb81c0c54de7c6943fcdf479c&wlanacname=3d1cd94ffbf7e4197e8fbd46a5584e53&ssid=&nasip=39ac2c6e007df760ae8b3f7f3b919dfe&snmpagentip=&mac=4e322ca419aeaaa5523942da438b26de&t=wireless-v2&url=709db9dc9ce334aa6363270493a5e6a6b1748319c9795b5e&apmac=&nasid=3d1cd94ffbf7e4197e8fbd46a5584e53&vid=1b33d3067b548968&port=2b0765f54b94f6f7&nasportid=5b9da5b08a53a54010ce97b909267f4e49b8dcf9acf28fa02ad8591e2fe4335e")
	query.Add("operatorPwd", "")
	query.Add("operatorUserId", "")
	query.Add("validcode", "")
	req.URL.RawQuery = query.Encode()
	// 设置header
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36")
	req.Header.Set("Referer", "http://210.34.48.49/eportal/index.jsp?wlanuserip=4168a23cb81c0c54de7c6943fcdf479c&wlanacname=3d1cd94ffbf7e4197e8fbd46a5584e53&ssid=&nasip=39ac2c6e007df760ae8b3f7f3b919dfe&snmpagentip=&mac=4e322ca419aeaaa5523942da438b26de&t=wireless-v2&url=709db9dc9ce334aa6363270493a5e6a6b1748319c9795b5e&apmac=&nasid=3d1cd94ffbf7e4197e8fbd46a5584e53&vid=1b33d3067b548968&port=2b0765f54b94f6f7&nasportid=5b9da5b08a53a54010ce97b909267f4e49b8dcf9acf28fa02ad8591e2fe4335e")
	res, err := client.Do(req)
	if err != nil {
		return *new(LoginResponse)
	}
	defer res.Body.Close()
	loginRes := new(LoginResponse)
	_ = json.NewDecoder(res.Body).Decode(loginRes)
	return *loginRes
}
