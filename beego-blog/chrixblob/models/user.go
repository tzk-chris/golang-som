package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int        `orm:"pk;auto"`
	UserName   string     `orm:"description(用户名);index;unique"`
	Password   string     `orm:"description(密码)"`
	IsAdmin    int        `orm:"description(1:管理员, 2:普通用户);default(2)"`
	CreateTime time.Time  `orm:"auto_now;type(datetime);description(创建时间)"`
	Cover      string     `orm:"description(用户头像);default(static/upload/bq2.png)"`
	Posts      []*Post    `orm:"reverse(many)"`
	Comments   []*Comment `orm:"reverse(many)"`
}

func (u *User) TableName() string {
	return "sys_user"
}

func init() {
	orm.RegisterModel(new(User), new(Post), new(Comment))
}
