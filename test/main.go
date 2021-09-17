package main

import (
	"fmt"
	"github.com/esourceful/xlsxreader"
	"database/sql"
	//"time"
	_ "github.com/go-sql-driver/mysql"
	/*	"github.com/esourceful/xlsxreader"*/
	"encoding/csv"
	"log"
	"os"
)

func dbSql(dsn string, query string){
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/belfserv_B")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	else {
		Println("database connection was successful!")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}


	//charset=utf8mb4,utf8
}

func main() {
	fmt.Println("starting test/main.go")
	filepath := "/Users/johndavis/go/src/r/xlsxreader/test/sample_parse_failure.xlsx"
	//filepath := "./test-small.xlsx"
	e, err := xlsxreader.OpenFile(filepath)
	if err != nil {
		fmt.Printf("error: %s \n", err)
		return
	}
	defer e.Close()

	fmt.Printf("Worksheets: %s \n", e.Sheets)

	rowcount := 0
	cellcount := 0
	columncount := 0
	brokencount := 0


	for row := range e.ReadRows(e.Sheets[0]) {
		fmt.Println("reading row")
		if row.Error != nil {
			fmt.Printf("error on row %d: %s \n", row.Index, row.Error)
			return
		}
		/*
			if row.Index < 10 {
				fmt.Printf("%+v \n", row.Cells)
			}*/
		rowcount++

		record := make([]string, 0, len(row.Cells))

		for _, cell := range row.Cells {
			cellcount++
			record = append(record, cell.Value)
			if rowcount == 1 {
				columncount++
				fmt.Println("found cell header value:", cell.Value)
			}
		}
		if len(record) != columncount {
			brokencount++
			fmt.Println(rowcount, ": found row that has count different from columncount (", len(record), "vs", columncount, ")")
			for i := 0; i < len(record); i++ {
				//fmt.Println(i, record[i])
			}
			/*if brokencount > 2 {
				break
			}*/
		}
	}
	fmt.Println("there were",brokencount,"broken records found out of",rowcount,"records !")
}

func main2() {
	filepath := "./sample_parse_failure.xlsx"
	//filepath := "./test-small.xlsx"
	e, err := xlsxreader.OpenFile(filepath)
	if err != nil {
		fmt.Printf("error: %s \n", err)
		return
	}
	defer e.Close()

	fmt.Printf("Worksheets: %s \n", e.Sheets)

	rowcount := 0
	cellcount := 0
	columncount := 0
	brokencount := 0


	for row := range e.ReadRows(e.Sheets[0]) {
		/*
		if row.Error != nil {
			fmt.Printf("error on row %d: %s \n", row.Index, row.Error)
			return
		}

		if row.Index < 10 {
			fmt.Printf("%+v \n", row.Cells)
		}*/
		rowcount++

		record := make([]string, 0, len(row.Cells))

		for _, cell := range row.Cells {
			cellcount++
			record = append(record, cell.Value)
			if rowcount == 1 {
				columncount++
				fmt.Println("found cell header value:", cell.Value)
			}
		}
		if len(record) != columncount {
			brokencount++
			//fmt.Println(rowcount, ": found row that has count different from columncount (", len(record), "vs", columncount, ")")
			for i := 0; i < len(record); i++ {
				//fmt.Println(i, record[i])
			}
			/*if brokencount > 2 {
				break
			}*/
		}
	}
}
