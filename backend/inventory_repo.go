package main

import (
	"log"
	"net/url"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetInventory(id int) Inventory {
	var i Inventory
	row := GetDB().QueryRow("SELECT id,book_id,guest_id,status,last_updated FROM Inventory WHERE id=?", id)
	err := row.Scan(&i.ID, &i.BookID, &i.GuestID, &i.Status, &i.LastUpdated)
	if err != nil {
		log.Println(err.Error())
	}
	return i
}

func GetInventoryList() []Inventory {
	rows, err := GetDB().Query("SELECT id,book_id,guest_id,status,last_updated FROM Inventory")
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	defer rows.Close()

	var InventoryList []Inventory
	for rows.Next() {
		var i Inventory
		err := rows.Scan(&i.ID, &i.BookID, &i.GuestID, &i.Status, &i.LastUpdated)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		InventoryList = append(InventoryList, i)
	}
	err = rows.Err()
	if err != nil {
		log.Println(err.Error())
	}
	return InventoryList
}

func CreateInventory(form url.Values) error {
	insert, err := GetDB().Query("INSERT INTO Inventory (book_id) VALUES (?)",
		form.Get("book_id"))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	insert.Close()
	return nil
}

func EditInventory(id int, form url.Values) error {
	update, err := GetDB().Query("UPDATE Inventory SET book_id=?,guest_id=?,status=?,last_updated=? WHERE id=?",
		form.Get("book_id"), form.Get("guest_id"), form.Get("status"), time.Now(), form.Get("description"), id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func CheckoutInventory(id int, guestID int) error {
	update, err := GetDB().Query("UPDATE Inventory SET guest_id=?,status=?,last_updated=? WHERE id=?",
		guestID, "out", time.Now(), id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func ReturnInventory(id int) error {
	update, err := GetDB().Query("UPDATE Inventory SET status=?,last_updated=? WHERE id=?",
		"in", time.Now(), id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	update.Close()
	return nil
}

func DeleteInventory(id int) error {
	delete, err := GetDB().Query("DELETE FROM Inventory WHERE id=?", id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	delete.Close()
	return nil
}
