package main

import (
	"database/sql"
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func IsAdmin(cardID int) bool {
	var result int
	err := GetDB().QueryRow("SELECT card_id FROM Users WHERE card_id=? and role=1", cardID).Scan(&result)
	if err == nil {
		return true
	}
	if err == sql.ErrNoRows {
		return false
	}
	log.Println(err.Error())
	return false
}

func SetLastLogin(cardID int) {
	update, err := GetDB().Query("UPDATE Users SET last_login=? WHERE card_id=?", time.Now(), cardID)
	if err != nil {
		log.Println(err.Error())
	}
	update.Close()
}

func GetUser(id int) User {
	var u User
	row := GetDB().QueryRow("SELECT id,name,email,phone,card_id,role,last_login FROM Users WHERE id=?", id)
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.CardID, &u.Role, &u.LastLogin)
	if err != nil {
		log.Println(err.Error())
	}
	return u
}

func GetUsers() []User {
	rows, err := GetDB().Query("SELECT id,name,email,phone,card_id,role,last_login FROM Users")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var userList []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.CardID, &u.Role, &u.LastLogin)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		userList = append(userList, u)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return userList
}

func CreateUser(form url.Values) error {
	insert, err := GetDB().Query("INSERT INTO Users (name,email,phone,card_id,role) VALUES (?,?,?,?,?)",
		form.Get("name"), form.Get("email"), form.Get("phone"), form.Get("card_id"), form.Get("role"))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func EditUser(id int, form url.Values) error {
	update, err := GetDB().Query("UPDATE Users SET name=?,email=?,phone=?,card_id=?,role=? WHERE id=?",
		form.Get("name"), form.Get("email"), form.Get("phone"), form.Get("card_id"), form.Get("role"), id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func DeleteUser(id int) error {
	delete, err := GetDB().Query("DELETE FROM Users WHERE id=?", id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}
