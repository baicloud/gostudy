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
	Id   int64  `db:"id"`
	Name string `db:"string"`
	Age  int    `db:"age"`
}

func testQueryData() {
	for i := 0; i < 150; i++ {
		fmt.Printf("query %d times\n", i)
		sqlstr := "select id,name,age from user where id=?"
		row := DB.QueryRow(sqlstr, 1)
		var user User
		// row必须scan,不然会导致连接无法关闭，会一直占用连接
		continue
		err := row.Scan(&user.Id, &user.Name, &user.Age)
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%v age:%d\n", user.Id, user.Name, user.Age)
	}
}

func testQueryMulRow() {
	sqlstr := "select id,name,age from user where id>?"
	rows, err := DB.Query(sqlstr, 0)
	//重点，rows对象一定要close掉
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
		err := rows.Scan(&user.Id, &user.Name, &user.Age) //不scan会导致连接不释放
		if err != nil {
			fmt.Printf("scan failed,err:%v\n", err)
			return
		}
		fmt.Printf("user:%#v\n", user)
	}
}

func testInsertData() {
	sqlstr := "insert into user(name,age) values(?,?)"
	result, err := DB.Exec(sqlstr, "tom", 10)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Printf("get last insert id failed,err:%v\n", err)
		return
	}
	fmt.Printf("id is %d\n", id)
}

func testUpdateData() {
	sqlstr := "update user set name=? where id=?"
	result, err := DB.Exec(sqlstr, "jim", 3)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed,err:%v\n", err)
		return
	}
	fmt.Printf("update db succ,affected rows:%d\n", affected)
}

func testDeleteData() {
	sqlstr := "delete from user where id=?"
	result, err := DB.Exec(sqlstr, 3)
	if err != nil {
		fmt.Printf("insert failed,err:%v\n", err)
		return
	}
	affected, err := result.RowsAffected()
	if err != nil {
		fmt.Printf("get affected rows failed,err:%v\n", err)
		return
	}
	fmt.Printf("delete db succ,affected rows:%d\n", affected)
}

func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// testQueryData()
	// testQueryMulRow()
	// testInsertData()
	// testUpdateData()
	testDeleteData()
}
