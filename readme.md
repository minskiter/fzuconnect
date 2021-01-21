### FZU 福州大学校园网自动连接服务

#### 配置文件 config.ini
``` ini
[common]
username=
password=

[service]
name = FZUConnect
displayname = FZU校园网自动连接服务
description = 福州大学校园网自动登陆服务
```

#### 注册服务
``` sh
./fzuconnect install
```

自定义配置文件路径
``` sh
./fzuconnect install -c ./config.ini 
```

#### 开始服务
``` sh
./fzuconnect start
```

### 停止服务
``` sh
./fzuconnect stop
```

### 重启服务
``` sh
./fzuconnect restart
```

### 日志查看
window 打开事件查看器
```
win+r eventvwr
```
使用来源FzuConnect过滤日志即可查看完整日志