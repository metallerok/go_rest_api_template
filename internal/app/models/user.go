package models

import (
	"database/sql"
	"log"
)

type User struct {
	ID    int64  `json:"id"`
	Type  string `json:"type"`
	Token string `json:"token"`
}

func UserGetByToken(token string) int64 {
	var uid int64
	err := db.LocalDB.QueryRow(
		`SELECT 
			id 
		FROM 
			user 
		WHERE 
			token = $1 AND type = 'ADMIN'
		`,
		token,
	).Scan(&uid)
	if err == sql.ErrNoRows {
		return 0
	} else if err != nil {
		log.Fatal(err)
	}
	return uid
}
