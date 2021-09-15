package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	id int `db:"id"`
	name string `db:"name"`
	age int `db:"age"`
	phone string `db:"phone"`
}

var db *sqlx.DB

//初始化数据库连接，init()方法会在系统调用main()方法前执行
func init() {
	database, err := sqlx.Open("mysql", "root:105293@tcp(localhost:3306)/nba")
	if err != nil {
		fmt.Println("open mysql failed,", err)
	}
	db = database
}

func main() {
	//添加用户
	sql := "insert into user(name, age, phone) values(?, ?, ?)"
	ret, err := db.Exec(sql, "mzl", 19, "13652707142")
	if err != nil{
		fmt.Println("insert fail: ", err)
		return
	}

	//查询最后一个用户
	id, err := ret.LastInsertId()
	if err != nil{
		fmt.Println("select last fail: ", err)
		return
	}
	fmt.Println("insert success: ", id, "  ", ret)

}