package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func initDb() error {
	var err error
	dsn := "root:root@tcp(localhost:3306)/golang_db"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	DB.SetConnMaxLifetime(100 * time.Second)
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}

type User struct {
	Id   int64          `db:"id"`
	Name sql.NullString `db:"string"`
	Age  int            `db:"age"`
}

func testPrepareData() {
	sqlstr := "select id, name , age from user where id > ?"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("prepare failed,err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()

	rows, err := stmt.Query(0)

	//rows一点要close掉
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	if err != nil {
		fmt.Printf("query failed,err:%v\n", err)
		return
	}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("user:%v\n", user)
	}
}

func testPrepareInsertData() {
	sqlstr := "insert into user(name,age) value(?,?)"
	stmt, err := DB.Prepare(sqlstr)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	defer func() {
		if stmt != nil {
			stmt.Close()
		}
	}()
	result, err := stmt.Exec("jim", 100)
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed,err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed,err:%v", err)
		return
	}
	// testPrepareData()
	testPrepareInsertData()
}
