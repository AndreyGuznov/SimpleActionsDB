package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string
	Age  int
}

func main() {
	db, err := sql.Open("mysql", "root:PASSWORD@/forweb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	insert, err := db.Query("INSERT INTO `user` (`name`,`age`) VALUES ('Lens', 30)")
	if err != nil {
		panic(err)
	}
	defer insert.Close()

	res, err := db.Query("SELECT `name`,`age` FROM `user`")
	if err != nil {
		panic(err)
	}
	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Age)
		if err != nil {
			panic(err)
		}
		a := fmt.Sprintf("User: %s with age %d", user.Name, user.Age)
		fmt.Println(a)
	}
}
