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

	// insertRow(db, &employees{Name: "张三", Department: "技术部", Salary: 3000})
	// insertRow(db, &employees{Name: "李四", Department: "技术部", Salary: 3500})
	// insertRow(db, &employees{Name: "王五", Department: "市场部", Salary: 3200})

	queryRow(db)
}

type employee struct {
	ID         uint64  `db:"id"`
	Name       string  `db:"name"`
	Department string  `db:"department"`
	Salary     float64 `db:"salary"`
}

func createTable(db *sqlx.DB) error {
	query := `
		create table if not exists employees (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			name VARCHAR(128) NOT NULL,
			department VARCHAR(128) NOT NULL,
			salary DOUBLE NOT NULL
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("ERROR:创建数据表失败: %v", err)
		return err
	}

	return err
}

func insertRow(db *sqlx.DB, employee *employee) error {
	_, err := db.Exec(
		"insert into employees (name, department, salary) values (?, ?, ?)",
		employee.Name, employee.Department, employee.Salary,
	)
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}

	return err
}

func queryRow(db *sqlx.DB) error {
	var employees []employee
	err := db.Select(&employees, "select * from employees where department = ?", "技术部")
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}

	fmt.Printf("查询结果:%v", employees)

	err = db.Select(&employees, `
		select * from employees 
		where salary = (select max(salary) from employees)
	`)
	if err != nil {
		log.Fatal("查询数据失败:", err)
	}

	fmt.Println()
	fmt.Printf("查询结果:%v", employees)

	return err
}
