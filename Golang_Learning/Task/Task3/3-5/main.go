package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//连接数据库
	dsn := "root:test@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}

	//创建表
	createTable(db)

	//insertRow(db)

	//searchRow(db)

	//searchRow2(db)

	deleteRow(db)
}

type User struct {
	gorm.Model
	Name    string
	PostNum uint32
	Posts   []Post `gorm:foreignKey:UserID`
}

type Post struct {
	gorm.Model
	Title    string
	UserID   uint
	State    string
	Comments []Comment
}

func (post *Post) AfterCreate(tx *gorm.DB) (err error) {
	tx.Model(&User{}).Where("id = ?", post.UserID).
		Update("post_num", gorm.Expr("post_num + 1"))

	return nil
}

type Comment struct {
	gorm.Model
	PostID  uint
	Content string
}

func (comment *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var commentCount uint64
	tx.Model(&Comment{}).Where("post_id = ?", comment.PostID).
		Select("count(id) as comment_count").Scan(&commentCount)
	if commentCount == 0 {
		tx.Model(&Post{}).Where("id = ?", comment.PostID).
			Update("state", "无评论")
	}

	return nil
}

func createTable(db *gorm.DB) error {
	err := db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Printf("ERROR:创建数据表失败: %v", err)
		return err
	}

	return err
}

func insertRow(db *gorm.DB) {
	users := []User{
		{Name: "张三", Posts: []Post{
			{Title: "张三的文章", Comments: []Comment{
				{Content: "写的真好"},
			}},
			{Title: "张三的文章2", Comments: []Comment{
				{Content: "写的真棒"},
				{Content: "厉害"},
				{Content: "哈哈"},
			}},
		}},
		{Name: "李四", Posts: []Post{
			{Title: "李四的文章", Comments: []Comment{
				{Content: "写的真6"},
				{Content: "666"},
			}},
		}},
	}

	db.Create(&users)
}

func searchRow(db *gorm.DB) {
	var users []User
	db.Model(&User{}).Preload("Posts").Preload("Posts.Comments").
		Where("name = ?", "张三").Find(&users)
	fmt.Printf("查询结果:%v", users)
}

func searchRow2(db *gorm.DB) {
	var post Post
	db.Model(&Post{}).
		Joins("left join comments on posts.id = comments.post_id").
		Select("posts.*, count(posts.id) as comment_count").
		Group("posts.id").Order("comment_count DESC").
		First(&post)
	fmt.Printf("查询结果:%+v", post)
}

func deleteRow(db *gorm.DB) {
	comments := []Comment{}
	db.Model(&Comment{}).Where("id in (?, ?)", 5, 6).Find(&comments)
	db.Delete(&comments)
}
