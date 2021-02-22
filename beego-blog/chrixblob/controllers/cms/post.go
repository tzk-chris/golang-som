package cms

import (
	"chrixblob/models"
	"chrixblob/tools"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
	"time"
)

type PostController struct {
	beego.Controller
}

func (p *PostController) Get() {

	o := orm.NewOrm()
	posts := []models.Post{}

	qs := o.QueryTable(new(models.Post)).RelatedSel()
	qs.All(&posts) // RelatedSel() -- 是通过外键的关联查询到User表中的UserName用到的

	count, _ := qs.Count() // 获取数据库中表中总共有几条数据
	current_page, err := p.GetInt("page")
	if err != nil {
		current_page = 1
	}
	page_size := 2

	total_page := tools.GetPageNum(count, page_size)
	// 前后页码
	aroud_count := 1 // 设置显示前后页码数
	left_pages, right_pages, left_has_more, right_has_more := tools.Get_pagination_data(total_page, current_page, aroud_count)
	has_pre_page, has_next_page, pre_page, next_page := tools.HasNext(current_page, total_page)

	p.Data["left_pages"] = left_pages
	p.Data["left_has_more"] = left_has_more
	p.Data["current_page"] = current_page

	p.Data["has_pre_page"] = has_pre_page
	p.Data["pre_page"] = pre_page
	p.Data["has_next_page"] = has_next_page
	p.Data["next_page"] = next_page

	p.Data["right_pages"] = right_pages
	p.Data["right_has_more"] = right_has_more

	p.Data["num_pages"] = total_page //总页数
	p.Data["count"] = count          //总数量
	p.Data["posts"] = posts
	p.TplName = "cms/post-list.html"
}

func (p *PostController) ToAdd() {
	p.TplName = "cms/post-add.html"
}
func (p *PostController) DoAdd() {

	title := p.GetString("title")
	description := p.GetString("description")
	content := p.GetString("content")

	f, h, err := p.GetFile("cover")

	defer f.Close()
	var cover string
	if err != nil {
		cover = "static/upload/no_pic.jpg"
	}
	// 生成时间戳，防止重名
	timeUnix := time.Now().Unix()               // int64类型
	time_str := strconv.FormatInt(timeUnix, 10) // 将int64转为字符串 convert：转换

	path := "static/upload/" + time_str + h.Filename
	// 保存获取到的文件
	err1 := p.SaveToFile("cover", path)

	if err1 != nil {
		cover = "static/upload/no_pic.jpg"
	}
	cover = path

	// 添加到数据库
	o := orm.NewOrm()
	author := p.GetSession("cms_user_name")
	user := models.User{}
	o.QueryTable(new(models.User)).Filter("user_name", author).One(&user)
	post := models.Post{
		Title:       title,
		Description: description,
		Content:     content,
		Cover:       cover,
		Author:      &user,
	}

	_, err2 := o.Insert(&post)

	if err2 != nil {
		fmt.Println("=============")
		fmt.Println(err2)
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": err2}
		p.ServeJSON()
	}

	p.Data["json"] = map[string]interface{}{"code": 200, "msg": "添加成功"}
	p.ServeJSON()

}

func (p *PostController) PostDelete() {

	id, err := p.GetInt("id")
	if err != nil {
		p.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()
	_, err2 := o.QueryTable(new(models.Post)).Filter("id", id).Delete()

	if err2 != nil {
		fmt.Println(err2)
		p.Ctx.WriteString("删除错误")
	}

	p.Redirect(beego.URLFor("PostController.Get"), 302)

}

func (p *PostController) ToEdit() {

	id, err := p.GetInt("id")
	if err != nil {
		p.Ctx.WriteString("id参数错误")
	}

	o := orm.NewOrm()

	post := models.Post{}
	o.QueryTable(new(models.Post)).Filter("id", id).One(&post)
	p.Data["post"] = post
	p.TplName = "cms/post-edit.html"

}

func (p *PostController) DoEdit() {

	o := orm.NewOrm()

	id, err := p.GetInt("id")
	if err != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": "id参数错误"}
	}

	qs := o.QueryTable(new(models.Post)).Filter("id", id)

	title := p.GetString("title")
	description := p.GetString("description")
	content := p.GetString("content")

	f, h, err1 := p.GetFile("cover")

	//fmt.Println("==============")
	//fmt.Println(id)
	//fmt.Println(title)
	//fmt.Println(description)
	//fmt.Println(content)
	//fmt.Println(err1)

	if err1 != nil {

		_, err4 := qs.Update(orm.Params{
			"title":   title,
			"description":    description,
			"content": content,
		})

		if err4 != nil {
			p.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
		}
		p.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
		p.ServeJSON()

	}

	defer f.Close()

	// 生成时间戳，防止重名
	timeUnix := time.Now().Unix()               // int64类型
	time_str := strconv.FormatInt(timeUnix, 10) // 将int64转为字符串 convert：转换

	path := "static/upload/" + time_str + h.Filename
	// 保存获取到的文件
	err2 := p.SaveToFile("cover", path)

	if err2 != nil {
		_, err5 := qs.Update(orm.Params{
			"title":   title,
			"description":    description,
			"content": content,
		})
		if err5 != nil {
			p.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
		}
		p.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
		p.ServeJSON()
	}
	_, err6 := qs.Update(orm.Params{
		"title":   title,
		"description":    description,
		"content": content,
		"cover":   path,
	})

	if err6 != nil {
		p.Data["json"] = map[string]interface{}{"code": 500, "msg": "更新失败"}
	}
	p.Data["json"] = map[string]interface{}{"code": 200, "msg": "更新成功"}
	p.ServeJSON()

}
