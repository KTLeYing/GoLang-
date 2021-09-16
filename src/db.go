package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Age   int    `db:"age"`
	Phone string `db:"phone"`
}

type User1 struct {
	id    int    `db:"id"`
	name  string `db:"name"`
	age   int    `db:"age"`
	phone string `db:"phone"`
}

var db *sqlx.DB
var index int64

//初始化数据库连接，init()方法会在系统调用main()方法前执行
func init() {
	database, err := sqlx.Open("mysql", "root:105293@tcp(localhost:3306)/nba")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}
	db = database
}

//添加用户
func add() {
	sql := "insert into user(name, age, phone) values(?, ?, ?)"
	res, err := db.Exec(sql, "v_ktlema", 19, "13652707142")
	if err != nil {
		fmt.Println("insert fail: ", err)
		return
	}

	//查询最后一个用户
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("select last fail: ", err)
		return
	}
	index = id
	fmt.Println("insert success: ", id, "  resultSet: ", res)
}

//删除用户
func delete() {
	sql := "delete from user where id = ?"
	res, err := db.Exec(sql, 3)
	if err != nil {
		fmt.Println("delete fail: ", err)
		return
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row affect fail", err)
		return
	}
	fmt.Println("delete success", row)
}

//更新用户
func update() {
	sql := "update user set age = 20 where id = ?"
	res, err := db.Exec(sql, 5)
	if err != nil {
		fmt.Println("update fail: ", err)
	}

	row, err := res.RowsAffected()
	if err != nil {
		fmt.Println("row affect fail")
		return
	}
	fmt.Println("update success,", row)
}

//查询用户
func find() {
	var user []User
	sql := "select id, name, age, phone from user"
	err := db.Select(&user, sql)
	if err != nil {
		fmt.Println("select fail", err)
		return
	}
	fmt.Println("select success,", user)
}

//查询1
func find1() {
	sql := "select * from user"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("select fail", err)
		return
	}

	//遍历数据
	for rows.Next() {
		var id int
		var username string
		var age int
		var phone string
		err := rows.Scan(&id, &username, &age, &phone)
		if err != nil {
			fmt.Println("rows scan fail")
			return
		}
		fmt.Printf("id:%d  name:%s  age:%d  phone:%s\n", id, username, age, phone)
	}
}

//查询2
func find2() {
	sql := "select * from user"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("select fail", err)
		return
	}

	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	//遍历获取结果集数据
	for rows.Next() {
		var user User1
		err := rows.Scan(&user.id, &user.name, &user.age, &user.phone)
		if err != nil {
			fmt.Println("rows scan fail")
			return
		}
		fmt.Printf("id:%d  name:%s  age:%d  phone:%s\n", user.id, user.name, user.age, user.phone)
	}
}

func transaction() {
	var sql string
	//开始事务
	conn, err := db.Begin()
	if err != nil {
		fmt.Println("begin fail,", err)
		return
	}

	//执行事务
	//插入操作1
	sql = "insert into user(name, age, phone) values(?, ?, ?)"
	res, err := conn.Exec(sql, "lisi", 20, "13652707142")
	if err != nil {
		fmt.Println("insert fail,", err)
		return
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("row fail", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success, ", id)

	//插入操作2
	res1, err := conn.Exec(sql, "wangwu", 21, "13622844791")
	if err != nil {
		fmt.Printf("insert fail, ", err)
		conn.Rollback()
		return
	}
	id1, err := res1.LastInsertId()
	if err != nil {
		fmt.Println("row fail,", err)
		conn.Rollback()
		return
	}
	fmt.Println("insert success, ", id1)

	//提交事务
	conn.Commit()

}

func main() {
	//add()
	//delete()
	//update()
	//find()
	//find1()
	//find2()
	transaction()
}
