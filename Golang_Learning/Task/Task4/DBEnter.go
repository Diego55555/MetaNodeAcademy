package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var g_db *gorm.DB

type User struct {
	gorm.Model
	UserName string `gorm:"unique;not null"`
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
func insertRecordUser(user *User) (err error) {
	err = g_db.Create(user).Error
	if err != nil {
		log.Printf("[ERROR]创建用户失败: %v", err)
	}

	return
}

// 根据名称查找用户
func queryRecordUserByName(name string) (user User, err error) {
	err = g_db.Model(User{}).Where("user_name = ?", name).First(&user).Error
	if err != nil {
		log.Printf("[ERROR]查找用户失败: %v", err)
	}

	return
}

// 根据ID查找用户
func queryRecordUserByID(id uint) (user User, err error) {
	err = g_db.Model(User{}).Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("[ERROR]查找用户失败: %v", err)
	}

	return
}

// 创建文章
func insertRecordPost(post *Post) (err error) {
	err = g_db.Create(post).Error
	if err != nil {
		log.Printf("[ERROR]创建文章失败: %v", err)
	}

	return
}

// 根据文章ID查询文章信息
func queryRecordPostByID(id uint) (post Post, err error) {
	err = g_db.Model(Post{}).Where("id = ?", id).First(&post).Error
	if err != nil {
		log.Printf("[ERROR]查找文章失败: %v", err)
	}

	return
}

// 更具用户ID查询所有文章信息
func queryRecordPostByUserID(userID uint) (posts []Post, err error) {
	err = g_db.Model(Post{}).Where("user_id = ?", userID).Find(&posts).Error
	if err != nil {
		log.Printf("[ERROR]查找文章失败: %v", err)
	}

	return
}

// 根据文章ID修改文章信息
func updateRecordPost(id uint, updates map[string]interface{}) (err error) {
	err = g_db.Model(Post{}).Where("id = ?", id).Updates(updates).Error
	if err != nil {
		log.Printf("[ERROR]更新文章失败: %v", err)
	}

	return
}

// 根据文章ID查询文章信息
func deleteRecordPostByID(id uint) (err error) {
	err = g_db.Where("id = ?", id).Delete(&Post{}).Error
	if err != nil {
		log.Printf("[ERROR]删除文章失败: %v", err)
	}

	return
}

// 创建评论
func insertRecordComment(comment *Comment) (err error) {
	err = g_db.Create(comment).Error
	if err != nil {
		log.Printf("[ERROR]创建评论失败: %v", err)
	}

	return
}

// 更具文章ID查询所有评论信息
func queryRecordCommentByPostID(postID uint) (comments []Comment, err error) {
	err = g_db.Model(Comment{}).Where("post_id = ?", postID).Find(&comments).Error
	if err != nil {
		log.Printf("[ERROR]查找评论失败: %v", err)
	}

	return
}
