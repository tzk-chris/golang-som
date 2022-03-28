package controller

import (
	"fmt"
	"gin-blog/dao"
	"gin-blog/model"
	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	"html/template"
	"strconv"
)

//func AddUser(c *gin.Context) {
//	username := c.PostForm("username")
//	password := c.PostForm("password")
//
//	user := model.User{
//		Username: username,
//		Password: password,
//	}
//
//	dao.Mgr.AddUser(&user)
//
//	c.Redirect(200, "/")
//}
//
//func ListUser(c *gin.Context) {
//	c.HTML(200, "user.html", nil)
//}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.AddUser(&user)

	c.Redirect(301, "/")
}

func GoRegisterUser(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)
	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "用户名不存在！")
		fmt.Println("用户名不存在！")
	} else {
		if u.Password != password {
			fmt.Println("密码错误！")
		} else {
			fmt.Println("登录成功！")
			c.Redirect(301, "/")
		}
	}
}

func GoLoginUser(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context)  {
	s := c.Query("pid")
	fmt.Println(s)
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	// blackfriday 可以把md格式的转化为html的
	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title": p.Title,
		"Content": template.HTML(content),
	})
}
