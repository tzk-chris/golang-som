package routers

import (
	_ "chrixblob/controllers"
	"chrixblob/controllers/cms"
	"chrixblob/controllers/front"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	// cms
	beego.Router("/cms", &cms.LoginController{})
	beego.Router("/cms/main/main", &cms.MainController{})
	beego.Router("/cms/main/welcome", &cms.MainController{}, "get:Welcome")
	beego.Router("/cms/main/post", &cms.PostController{})
	beego.Router("/cms/main/post_to_add", &cms.PostController{}, "get:ToAdd")
	beego.Router("/cms/main/post_do_add", &cms.PostController{}, "post:DoAdd")
	beego.Router("/cms/main/post_delete", &cms.PostController{}, "get:PostDelete")
	beego.Router("/cms/main/post_to_edit", &cms.PostController{}, "get:ToEdit")
	beego.Router("/cms/main/post_do_edit", &cms.PostController{}, "post:DoEdit")


	//beego.Router("/front") 前端
	beego.Router("/", &front.IndexController{})
	beego.Router("/detail", &front.IndexController{}, "get:PostDetail")
	beego.Router("/register", &front.RegisterController{})
	beego.Router("/login", &front.FrontLoginController{})
	beego.Router("/comment", &front.CommentController{})

}
