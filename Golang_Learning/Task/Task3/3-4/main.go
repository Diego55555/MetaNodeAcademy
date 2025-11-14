package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	// 连接信息
	dsn := "root:test@tcp(127.0.0.1:3306)/go_learning?charset=utf8mb4&parseTime=True&loc=Local"

	// 打开数据库连接
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatal("打开数据库失败:", err)
	}
	defer db.Close() // 确保连接关闭

	err = createTable(db)
	if err != nil {
		return
	}

	// insertRow(db, &book{Title: "哈哈", Author: "张三", Price: 30})
	// insertRow(db, &book{Title: "呵呵", Author: "李四", Price: 50})
	// insertRow(db, &book{Title: "哇哇", Author: "王五", Price: 60})

	queryRow(db)
}

type book struct {
	ID     uint64  `db:"id"`
	Title  string  `db:"title"`
	Author string  `db:"author"`
	Price  float64 `db:"price"`
}

func createTable(db *sqlx.DB) error {
	query := `
		create table if not exists books (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			title VARCHAR(128) NOT NULL,
			author VARCHAR(128) NOT NULL,
			price DOUBLE NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("ERROR:创建数据表失败: %v", err)
		return err
	}

	return err
}

func insertRow(db *sqlx.DB, book *book) error {
	_, err := db.Exec(
		"insert into books (title, author, price) values (?, ?, ?)",
		book.Title, book.Author, book.Price,
	)
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}

	return err
}

func queryRow(db *sqlx.DB) error {
	var books []book
	err := db.Select(&books, "select * from books where price > ?", 50)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}

	fmt.Printf("查询结果:%v", books)

	return err
}
