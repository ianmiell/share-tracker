package main


import (
	"encoding/csv"
	"fmt"
    //"io/ioutil"
    "log"
	//"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)


type get_share interface {
	get_url() string
}


type share struct {
	code string
	current_price float64
}


func (s *share) get_url() string {
	var (
		cmdOut []byte
		err    error
	)
	cmdName := "bash"
	cmdArgs := []string{"-c", "curl -L 'https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=" + s.code + "&apikey=UA0A5BZ20QIH9DFM&datatype=csv' | sed -n 2p"}
	if cmdOut, err = exec.Command(cmdName, cmdArgs...).Output(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running curl command: ", cmdArgs)
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "We ran: ", cmdArgs)
	csvresp := string(cmdOut)
	fmt.Println(csvresp)
	return csvresp
}


func get_share_prices() {
	codes := get_codes()
	for i := 0; i < len(codes); i++ {
		time.Sleep(50000 * time.Millisecond)
		var code string = codes[i]
		fmt.Println(code)
		res := (&(share{code: code})).get_url()
		fmt.Printf("%s", res)
		r := csv.NewReader(strings.NewReader(res))
		record, err := r.Read()
		if err != nil {
			if record == nil {
				fmt.Println("Skipping code: " + code)
				continue
			}
		}
		price, err := strconv.ParseFloat(record[4],64)
		price *= 100
		fmt.Printf("%d", price)
		if err != nil { log.Fatal(err) }
		var price_in_pence int = int(price - 0.5)
		fmt.Println(price_in_pence)
		insert_share_price(code, price_in_pence)
    }
}

//historical https://stackoverflow.com/questions/28885775/yahoo-finance-quotes-api-and-historical-data-api
//https://stackoverflow.com/questions/10040954/alternative-to-google-finance-api
