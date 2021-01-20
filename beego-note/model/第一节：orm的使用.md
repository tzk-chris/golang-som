# 第一节：orm的使用

## 一、mysql安装

- 用户名和密码的配置

#### 二、navicat安装

#### 三、beego中使用

- **安装**：

- go get github.com/astaxie/beego/orm

- go get github.com/go-sql-driver/mysql

- ###### 链接数据库：

- orm.RegisterDriver(“mysql”, orm.DRMySQL)

- orm.RegisterDataBase(“default”, “mysql”, “用户名:密码@tcp(IP:端口号)/数据库?charset=utf8”, 30)

- 参数一：数据库的别名，用来在 ORM 中切换数据库使用

- 参数二：驱动名称

- 参数三：对应的链接字符串

- 参数四(可选)：设置最大空闲连接

- 或根据数据库别名设置：orm.SetMaxIdleConns(“default”, 30)

- 参数五(可选)：设置最大数据库连接

- 或根据数据库别名设置： orm.SetMaxOpenConns(“default”, 30)

- ###### 注册模型

- 在init函数中：orm.RegisterModel(new(Article))，只有注册了模型才可以使用

- 创建表

- 在数据库中操作创建表即可，后面会讲自动迁移表

- ###### 时区设置：

- orm.DefaultTimeLoc = time.UTC // 设置为 UTC 时间

- ORM 在进行 RegisterDataBase 的同时，会获取数据库使用的时区，然后在 time.Time 类型存取时做相应转换，以匹配时间系统，从而保证时间不会出错

- #### 示例代码

- ###### 首先创建好数据库和表，表对应结构体中的属性

- ###### app.conf链接信息配置：

```
dbhost = 127.0.0.1
dbuser = root
dbpassword = root
dbport = 3306
db = user
```

#### main.go中初始化链接：

- ###### orm.RegisterDriver(“mysql”, orm.DRMySQL)

- ###### orm.RegisterDataBase(“default”, “mysql”, “root:root@/orm_test?charset=utf8”)

```
"github.com/astaxie/beego/orm"
_ "github.com/go-sql-driver/mysql"


func init() {
// 读取配置信息
dbhost := beego.AppConfig.String("dbhost")
dbport := beego.AppConfig.String("dbport")
dbuser := beego.AppConfig.String("dbuser")
dbpassword := beego.AppConfig.String("dbpassword")
db := beego.AppConfig.String("db")

//注册mysql Driver
orm.RegisterDriver("mysql", orm.DRMySQL)
//构造conn连接
//用户名:密码@tcp(url地址)/数据库
conn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + db + "?charset=utf8"
//注册数据库连接
orm.RegisterDataBase("default", "mysql", conn)
fmt.Printf("数据库连接成功！%s\n", conn)
}
```

#### 模型结构体：

```
注意：结构体的字段名首字母必须大写
type Article struct {
Id int
Name string
Author string

}
func init(){
orm.RegisterModel(new(Article))
}
```

#### 增加数据：

```
o := orm.NewOrm()
o.Using("default") // default可以不用using

newStudent := Student{Name:"zhiliao",Age:18,Gender:"男"}

//新增数据
id, err :=o.Insert(&newStudent) // 返回添加的数据id以及错误信息
```

#### 

#### 注意：

- ORM 必须注册一个别名为 default 的数据库，作为默认使用