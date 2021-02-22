package cms

import (
	"chrixblob/models"
	"chrixblob/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type LoginController struct {
	beego.Controller
}

func (l *LoginController) Get() {
	l.TplName = "cms/login.html"
}

func (l *LoginController) Post() {

	username := l.GetString("username") // 特殊字符过滤
	password := l.GetString("password") // 长度校验

	md5_pwd := tools.GetMd5(password)

	o := orm.NewOrm()
	exist := o.QueryTable(new(models.User)).Filter("username", username).Filter("password", md5_pwd).Exist()

	if exist {

		l.SetSession("cms_user_name", username)

		l.Redirect(beego.URLFor("MainController.Get"), 302)
		//l.Ctx.WriteString("logined")
	} else {
		l.Redirect(beego.URLFor("LoginController.Get"), 302)
	}

}
