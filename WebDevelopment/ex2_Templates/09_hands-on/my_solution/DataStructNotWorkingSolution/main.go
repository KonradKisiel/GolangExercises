package main

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type DataTable struct {
	Date      []string
	Open      []string
	High      []string
	Low       []string
	Close     []string
	Volume    []string
	Adj_Close []string
}

var Date []string
var Open []string
var High []string
var Low []string
var Close []string
var Volume []string
var Adj_Close []string

func main() {

	//Date,Open,High,Low,Close,Volume,Adj Close

	file, err := os.Open("table.csv")
	if err != nil {
		log.Fatalln(err)
	}

	r := csv.NewReader(file)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		Date = append(Date, record[0])
		Open = append(Open, record[1])
		High = append(High, record[2])
		Low = append(Low, record[3])
		Close = append(Close, record[4])
		Volume = append(Volume, record[5])
		Adj_Close = append(Adj_Close, record[6])
	}
	Table := DataTable{
		Date:      Date,
		Open:      Open,
		High:      High,
		Low:       Low,
		Close:     Close,
		Volume:    Volume,
		Adj_Close: Adj_Close,
	}

	err = tpl.Execute(os.Stdout, Table)
	if err != nil {
		log.Fatalln(err)
	}

}
