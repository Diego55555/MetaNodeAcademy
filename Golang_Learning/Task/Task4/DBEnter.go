package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var g_db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
}

type Post struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"not null"`
	UserID  uint
	User    User
}

type Comment struct {
	gorm.Model
	Content string `gorm:"not null"`
	UserID  uint
	User    User
	PostID  uint
	Post    Post
}

// 连接数据库
func connectDB() error {
	//连接数据库
	dsn := "root:test@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	g_db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("[ERROR]连接数据库失败: %v", err)
	}

	//创建表
	err = createTable()

	return err
}

// 创建数据表
func createTable() error {
	err := g_db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Printf("[ERROR]创建数据表失败: %v", err)
		return err
	}

	return err
}

// 创建用户
func insertUser(user *User) (err error) {
	err = g_db.Create(user).Error
	if err != nil {
		log.Printf("[ERROR]创建用户失败: %v", err)
	}

	return
}

// 根据名称查找用户
func queryUserByName(name string) (user User, err error) {
	err = g_db.Model(User{}).Where("name = ?", name).First(user).Error
	if err != nil {
		log.Printf("[ERROR]查找用户失败: %v", err)
	}

	return
}
