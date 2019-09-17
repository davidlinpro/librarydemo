package main

import (
	"log"
	"net/url"

	_ "github.com/go-sql-driver/mysql"
)

func GetBook(id int) Book {
	var b Book
	row := GetDB().QueryRow("SELECT id,title,author,isbn,description FROM Books WHERE id=?", id)
	err := row.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Description)
	if err != nil {
		log.Println(err.Error())
	}
	return b
}

func GetBooks() []Book {
	rows, err := GetDB().Query("SELECT id,title,author,isbn,description FROM Books")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var BookList []Book
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.ISBN, &b.Description)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		BookList = append(BookList, b)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return BookList
}

func CreateBook(form url.Values) error {
	insert, err := GetDB().Query("INSERT INTO Books (title,author,isbn,description) VALUES (?,?,?,?)",
		form.Get("title"), form.Get("author"), form.Get("isbn"), form.Get("description"))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func EditBook(id int, form url.Values) error {
	update, err := GetDB().Query("UPDATE Books SET title=?,author=?,isbn=?,description=? WHERE id=?",
		form.Get("title"), form.Get("author"), form.Get("isbn"), form.Get("description"), id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func DeleteBook(id int) error {
	delete, err := GetDB().Query("DELETE FROM Books WHERE id=?", id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}
