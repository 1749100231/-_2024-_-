package model

import "time"

type User struct {
	ID       int64  `json:"user_id"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"-"`
	UserType int    `gorm:"default:1"` //stu 1 admin 2
	//ArticleList string `json:"article_list" gorm:"default:'[]'"` //JSON 字符串
	CreatedAt time.Time
	UpdatedAt time.Time
}
