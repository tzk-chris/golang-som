package front

import (
	"chrixblob/models"
	"chrixblob/tools"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type FrontLoginController struct {
	beego.Controller
}

func (f *FrontLoginController) Get()  {
	f.TplName = "front/login.html"
}

func (f *FrontLoginController) Post() {

	username := f.GetString("username") // 特殊字符过滤
	password := f.GetString("password") // 长度校验

	md5_pwd := tools.GetMd5(password)

	o := orm.NewOrm()
	exist := o.QueryTable(new(models.User)).Filter("username", username).Filter("password", md5_pwd).Exist()

	if exist {

		f.SetSession("front_username", username)


		f.Redirect(beego.URLFor("IndexController.Get"), 302)
		//l.Ctx.WriteString("logined")
	} else {
		f.Redirect(beego.URLFor("FrontLoginController.Get"), 302)
	}

}
