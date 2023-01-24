# LeigodAutoPause
自动暂停雷神加速器

## 用法
下载 [Releases](https://github.com/ZeroY-Code/LeigodAutoPause/releases)  
修改config.ini中的username和password为你的雷神加速器账号  
programPath为雷神加速器程序位置  
运行程序会自动启动雷神加速器

```ini
# config.ini

#游戏进程列表文件位置
listPath = ./gamelist.txt
# 首次启动等待时间
firstStartWaitingTime = 60  
# 游戏退出后等待时间
gameExitWaitingTime = 120

[leigod]
username = 17700000000
password = 123456789
programPath = C:\Program Files (x86)\LeiGod_Acc\leigod.exe
```
```text
# gamelist.txt
steam.exe
LeagueClient.exe
```

## 运行机制
打开程序  
检测雷神加速器未启动则启动雷神加速器  
等待60秒(可修改firstStartWaitingTime), 此60秒为等待启动游戏时间  
开始检测游戏进程(gamelist.txt)  
检测到gamelist.txt的进程无一存活  
等待120秒(可修改gameExitWaitingTime), 此120秒为等待用户切换游戏时间  
如若120秒内用户未打开游戏, 则自动暂停加速器并关闭本程序


## 自行编译
```shell
go build -ldflags "-H=windowsgui"
```