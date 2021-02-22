package models

import "time"

type CommentTree struct {
	Id int
	Content string
	Author *User
	CreateTime time.Time
	Children []*CommentTree
}

type Comment struct {
	Id         int       `orm:"pk;auto"`
	Content    string    `orm:"size(4000);description(评论内容)"`
	Post       *Post     `orm:"rel(fk);description(帖子外键)"`
	Author     *User     `orm:"rel(fk);description(评论人)"`
	CreateTime time.Time `orm:"auto_now;type(datetime);description(创建时间)"`

	// 评论的父级id
	Pid int `orm:"default(0);description(父级评论)"`
}

func (c *Comment) TableName() string {
	return "sys_post_comment"
}
