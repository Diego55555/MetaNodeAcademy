package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 连接信息
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

	// err = insertRow(db)
	// if err != nil {
	// 	return
	// }

	// err = searchRow(db)
	// if err != nil {
	// 	return
	// }

	// err = updateRow(db)
	// if err != nil {
	// 	return
	// }

	// err = deleteRow(db)
	// if err != nil {
	// 	return
	// }

	handleTransaction(db, 1, 3, 100)
}

type account struct {
	id      uint64
	balance float64
}

type transaction struct {
	id            uint64
	fromAccountId uint64
	toAccountId   uint64
	amount        float64
}

func createTable(db *sql.DB) error {
	query := `
		create table if not exists accounts (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			balance DOUBLE
		)
	`
	_, err := db.Exec(query)
	if err != nil {
		log.Printf("ERROR:创建数据表失败: %v", err)
		return err
	}

	query = `
		create table if not exists transactions (
			id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
			from_account_id BIGINT UNSIGNED,
			to_account_id BIGINT UNSIGNED,
			amount DOUBLE
		)
	`
	_, err = db.Exec(query)
	if err != nil {
		log.Printf("ERROR:创建数据表失败: %v", err)
	}

	return err
}

func insertRow(db *sql.DB) error {
	// 插入数据
	_, err := db.Exec(
		"insert into accounts (balance) values (100)")
	if err != nil {
		log.Fatal("插入数据失败:", err)
	}

	return err
}

func handleTransaction(db *sql.DB, accountIdA uint64, accountIdB uint64, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("开启事务失败:", err)
	}

	result, err := tx.Exec(`
		update accounts set balance = balance - ? 
		where id = ? and balance >= ?
	`, amount, accountIdA, amount)
	if err != nil {
		log.Fatal("更新账户A失败:", err)
		tx.Rollback()
		return err
	}
	affected, _ := result.RowsAffected()
	if affected == 0 {
		err = errors.New("账户A ID不存在或者余额不足")
		log.Fatal("更新账户A失败:", err)
		tx.Rollback()
		return err
	}

	result, err = tx.Exec(`
		update accounts set balance = balance + ? 
		where id = ?
	`, amount, accountIdB)
	if err != nil {
		log.Fatal("更新账户B失败:", err)
		tx.Rollback()
		return err
	}
	affected, _ = result.RowsAffected()
	if affected == 0 {
		err = errors.New("账户B ID不存在或者余额不足")
		log.Fatal("更新账户B失败:", err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(`
		insert into transactions (from_account_id, to_account_id, amount) 
		values (?, ?, ?)
	`, accountIdA, accountIdB, amount)
	if err != nil {
		log.Fatal("添加交易记录失败:", err)
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
