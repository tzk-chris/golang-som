package front

import (
	"chrixblob/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type IndexController struct {
	beego.Controller
}

func (i *IndexController) Get() {

	o := orm.NewOrm()

	posts := []models.Post{}
	o.QueryTable(new(models.Post)).RelatedSel().All(&posts)

	front_username := i.GetSession("front_username")
	if front_username == nil {
		front_username = ""
	}

	i.Data["u"] = front_username
	fmt.Println(front_username)
	i.Data["posts"] = posts
	i.TplName = "front/index.html"

}

func (i *IndexController) PostDetail() {

	id, _ := i.GetInt("id")

	o := orm.NewOrm()
	post := models.Post{}
	qs := o.QueryTable(new(models.Post)).Filter("id", id)
	qs.RelatedSel().One(&post)

	// 阅读数+1
	qs.Update(orm.Params{"read_num": post.ReadNum + 1})

	front_username := i.GetSession("front_username")
	if front_username == nil {
		front_username = ""
	}

	// 构造评论树
	comments := []models.Comment{}
	o.QueryTable(new(models.Comment)).Filter("post_id", id).Filter("pid", 0).RelatedSel().All(&comments)
	comment_trees := []models.CommentTree{}
	for _, comment := range comments {
		pid := comment.Id
		comment_tree := models.CommentTree{
			Id:         comment.Id,
			Content:    comment.Content,
			Author:     comment.Author,
			CreateTime: comment.CreateTime,
			Children:   []*models.CommentTree{},
		}                            // 获取第一级的评论
		GetChild(pid, &comment_tree) // 用递归的方式获取每一层子级评论
		comment_trees = append(comment_trees, comment_tree)
	}

	i.Data["u"] = front_username
	i.Data["comment_trees"] = comment_trees
	i.Data["post"] = post
	i.TplName = "front/detail.html"

}

// 递归
func GetChild(pid int, comment_tree *models.CommentTree) {

	o := orm.NewOrm()
	qs := o.QueryTable(new(models.Comment))
	commments := []models.Comment{}
	_, err := qs.Filter("pid", pid).RelatedSel().All(&commments)
	if err != nil {
		return
	}

	// 查询二级及以下的多级评论
	for i := 0; i < len(commments); i++ {
		pid := commments[i].Id
		child := models.CommentTree{
			Id:         commments[i].Id,
			Content:    commments[i].Content,
			Author:     commments[i].Author,
			CreateTime: commments[i].CreateTime,
			Children:   []*models.CommentTree{},
		}
		comment_tree.Children = append(comment_tree.Children, &child)
		GetChild(pid, &child)
	}

	return
}

//[
//	{
//		id:1,
//		content:"xxx",
//		children:[
//			{
//				id:2,
//				content:"xxx",
//			},
//			{
//				id:3,
//				content:"xxx",
//			}
//		]
//
//	}
//]
