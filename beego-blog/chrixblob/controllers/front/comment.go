package front

import (
	"chrixblob/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type CommentController struct {
	beego.Controller
}

func (c CommentController) Post() {
	// 评论需要获取：1.评论文章的id 2.评论文章的评论内容 。。。
	// 获取评论文章的id
	post_id, err := c.GetInt("post_id")
	if err != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
		c.ServeJSON()
		return
	}

	o := orm.NewOrm()
	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id", post_id).One(&post)

	// 获取用户
	user_name := c.GetSession("front_username")
	if user_name == nil {
		c.Data["json"] = map[string]interface{}{"code": 401, "msg": "未登录"}
		c.ServeJSON()
		return
	}

	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", user_name).One(&user)

	// 获取评论内容
	content := c.GetString("content")
	pid, err1 := c.GetInt("pid")
	if err1 != nil {
		pid = 0
	}
	comment := models.Comment{
		Post:    &post,
		Content: content,
		Pid:     pid,
		Author:  &user,
	}

	_, err3 := o.Insert(&comment) // 插入评论

	if err3 != nil {
		c.Data["json"] = map[string]interface{}{"code": 500, "msg": "评论出错，请重试"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = map[string]interface{}{"code": 200, "msg": "评论成功"}
	c.ServeJSON()

}
