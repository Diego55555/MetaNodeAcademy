package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 链接信息
	dsn := "root:test@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() // 确保连接关闭

	err = createTable(db)
	if err != nil {
		return
	}

	err = insertRow(db)
	if err != nil {
		return
	}
}

type student struct {
	id    uint64
	name  string
	age   uint8
	grade string
}

func createTable(db *sql.DB) error {
	query := `
		create table if not exists students (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(128) NOT NULL,
			age TINYINT UNSIGNED,
			grade VARCHAR(32)
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("ERROR:打开数据库失败: %v", err)
	}

	return err
}

func insertRow(db *sql.DB) error {
	// 插入记录
	_, err := db.Exec(
		"insert into students (name, age, grade) values ('张三', 20, '三年级')")
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}

	return err
}
