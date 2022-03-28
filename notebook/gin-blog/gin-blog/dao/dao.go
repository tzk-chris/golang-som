package dao

import (
	"gin-blog/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User

	// 博客操作
	AddPost(post *model.Post)
	GetAllPost() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	db_conn := "username:password@tcp(ip:port)/databasename?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(db_conn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{
		db: db,
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

// 博客操作
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
func (mgr *manager) GetAllPost() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}
