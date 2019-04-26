package main

import (
	"time"
	"fmt"
	"os"
)

type transaction struct {
	code             string
	transaction_type string
	price            int
	number_of_shares int
	transaction_date time.Time
	notes            string
}



func create_transactions() []transaction {

	my_transactions := []transaction{}

    // Writing down the way the go standard time would look like formatted our way (2nd Jan)
	layout := "02/01/2006"
	date,_ := time.Parse(layout,"02/01/2006")
	// VARIOUS IN DATE ORDER DESC
	date, _ = time.Parse(layout,"28/11/2018")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 6390, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: "Ian"})
	date, _ = time.Parse(layout,"07/06/2018")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 13410, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: "Ian"})
	date, _ = time.Parse(layout,"30/11/2017")
	my_transactions = append(my_transactions,transaction{code: "WMH.L", price: 6390,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: "Ian"})
	date, _ = time.Parse(layout,"08/06/2017")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 12600, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"02/12/2016")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 6150,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"02/12/2016")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 6150,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"03/06/2016")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 12600, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"04/12/2015")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 6150,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"05/06/2015")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 12300, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"05/12/2014")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 6000,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"02/05/2012")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 10050, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"24/10/2012")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 5100,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"13/03/2013")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 11700, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"23/10/2013")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 5550,  number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"30/05/2014")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 11850, number_of_shares: 0, transaction_type: "DIVI", transaction_date: date, notes: ""})
	date, _ = time.Parse(layout,"19/04/2012")
	my_transactions = append(my_transactions,transaction{code: "WMH.L",  price: 26072,   number_of_shares: 1500, transaction_type: "BUY",  transaction_date: date, notes: ""})

	// Interest
	date, _ = time.Parse(layout,"20/08/2018")
	my_transactions = append(my_transactions,transaction{code: "INTEREST", price: 1564,    number_of_shares: 0, transaction_type: "INTEREST", transaction_date: date, notes: "Ian"})

	return my_transactions
}

func insert_transactions() {
	my_transactions := create_transactions()
	for i := 0; i < len(my_transactions); i++ {
		transaction_type := my_transactions[i].transaction_type
		code             := my_transactions[i].code
		price            := my_transactions[i].price
		number_of_shares := my_transactions[i].number_of_shares
		date             := my_transactions[i].transaction_date
		notes            := my_transactions[i].notes
        fmt.Fprintln(os.Stdout, "code: ", code)
        fmt.Fprintln(os.Stdout, "type: ", transaction_type)
		switch transaction_type {
			case "BUY":
				insert_share_tx(code, transaction_type, price, number_of_shares, date.UTC().Format(time.RFC3339))
			case "DIVI":
				insert_dividend(code, price, date.UTC().Format(time.RFC3339))
			case "SELL":
				insert_share_tx(code, transaction_type, price, number_of_shares, date.UTC().Format(time.RFC3339))
			case "INTEREST":
				insert_cash_tx("INTEREST", price, date.UTC().Format(time.RFC3339), notes)
		}
	}
}
