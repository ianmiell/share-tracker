package main

import (
	"database/sql"
	"log"
	//"time"
	_ "github.com/mattn/go-sqlite3"
)

func insert_share_price(code string, price int) {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("update share_price set price = ? where code = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(price, code)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	tx.Commit()
}


func insert_dividend(code string, dividend int, date string) {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or ignore into dividend(code, dividend, date) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(code, dividend, date)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	tx.Commit()
}


func insert_share_tx(code string, transaction_type string, price int, number_of_shares int, date string) {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert or ignore into share_tx(code, transaction_type, price, number_of_shares, date) values(?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(code, transaction_type, price, number_of_shares, date)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//if transaction_type == "BUY" {
	//	stmt, err = tx.Prepare("update share_price set num_held = num_held + ? where code = ?")
	//} else if transaction_type == "SELL" {
	//	stmt, err = tx.Prepare("update share_price set num_held = num_held - ? where code = ?")
	//}
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_, err = stmt.Exec(number_of_shares, code)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer stmt.Close()

	tx.Commit()
}


func insert_cash_tx(transaction_type string, price int, date string, notes string) {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("insert or ignore into cash_tx(transaction_type, price, date, notes) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(transaction_type, price, date, notes)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	tx.Commit()
}


func insert_share(code string, name string, notes string) {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert or ignore into share(code, name, notes) values(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(code, name, notes)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	tx.Commit()
}

func get_codes() []string {
	db, err := sql.Open("sqlite3", "./shares.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	rows, err := db.Query("select code from share")
	if err != nil {
	        log.Fatal(err)
	}
	defer rows.Close()
	codes := []string {}
	for rows.Next() {
		var code string
		err = rows.Scan(&code)
		if err != nil {
	        log.Fatal(err)
		}
		err = rows.Err()
		if err != nil {
		        log.Fatal(err)
		}
		codes=append(codes,code)
	}
	return codes
}
