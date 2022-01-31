package models

import (
	"log"
	"time"
)

var PlayedUser User

type User struct {
	ID int
	Name string
	Score int
	CreatedAt time.Time 
}

func (u *User) CreateUser() (err error) {  //Userを作成する
	cmd := `insert into users (
		name,
		score,
		created_at) values ($1, $2, $3)`

	_, err = Db.Exec(cmd,
		 u.Name, 
		 u.Score,
		 time.Now())

	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func GetTopUser() (topUsers []User) {  //Score上位10名を取得する
	cmd := `select * from users order by score desc limit 10`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		user := User{}
		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Score,
			&user.CreatedAt,
		)
		topUsers = append(topUsers, user)
	}
	return topUsers
}