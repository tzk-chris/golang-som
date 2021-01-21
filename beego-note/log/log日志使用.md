# 日志模块的使用

下载包：go get github.com/astaxie/beego/logs

## 一、使用

引入包：import “github.com/astaxie/beego/logs”

添加输出引擎：logs.SetLogger(“console”)

打日志：logs.Debug(“this is debug”)

###### SetLogger参数：

- 第一个参数：引擎

第二个参数：日志配置信息,字符串,不同的引擎参数不一样

```
{"filename":"project.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}
```

## 二、日志級別

1：LevelEmergency

2：LevelAlert

3：LevelCritical

4：LevelError

5：LevelWarn = LevelWarning

6：LevelNotice

7：LevelInfo = LevelInformational

8：LevelTrace = LevelDebug

## 三、引擎

1.console

logs.SetLogger(logs.AdapterConsole, `{"level":1,"color":true}`)

2.file

输出到文件

logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)

参数：

- filename 保存的文件名
- maxlines 每个文件保存的最大行数，默认值 1000000
- maxsize 每个文件保存的最大尺寸，默认值是 1 << 28, //256 MB 2^28 1<< 3 2^3
- daily 是否按照每天 logrotate，默认是 true
- maxdays 文件最多保存多少天，默认保存 7 天
- rotate 是否开启 logrotate，默认是 true
- level 日志保存的时候的级别，默认是 Trace 级别
- perm 日志文件权限 4(读权限)2(写权限)1(执行权限)

3.multifile

多文件日志写入，对号入座写入，比如test.error.log,err.debug.log

logs.SetLogger(logs.AdapterMultiFile, `{"filename":"test.log","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)

4.smtp

邮件发送

logs.SetLogger(logs.AdapterMail,`{"username":"xxx@qq.com","password":"认证密码","host":"smtp.qq.com:587","fromAddress":"xxx@qq.com","sendTos":["xxx@qq.com"]}`)

5…