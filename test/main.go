package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/esourceful/xlsxreader"
	_ "github.com/go-sql-driver/mysql"

	/*	"github.com/esourceful/xlsxreader"*/
	"encoding/csv"
	"log"
	"os"

	"runtime"
)

func dbSql(dsn string, query string) {
	db, err := sql.Open("mysql", "root:pass@tcp(127.0.0.1:3306)/belfserv_B")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("database connection was successful!")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	/*results, err := db.Query("SELECT id, name FROM tags")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	*/
	//charset=utf8mb4,utf8
}

func main() {
	start := time.Now()
	fmt.Println("starting test/main.go")
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	//get default filepath values
	fmt.Println(path)
	//filepath := "/Users/johndavis/go/src/r/xlsxreader/test/sample_parse_failure.xlsx"
	filepath := path + "/test/backlog.xlsx"
	//outfilepath := "/Users/johndavis/websites/GoWork/xlsxreader/sample.csv"
	outfilepath := path + "/test/sample_outfile.csv"

	//filepath := "./test-small.xlsx"
	e, err := xlsxreader.OpenFile(filepath)
	if err != nil {
		fmt.Printf("error: %s \n", err)
		return
	}
	defer e.Close()

	csvfile, err := os.Create(outfilepath)

	if err != nil {
		log.Fatalf("failed creating csv file: %s", err)
	}

	//open output csv file
	csvwriter := csv.NewWriter(csvfile)

	//todo: defer csvwriter close?

	fmt.Printf("Worksheets: %s \n", e.Sheets)

	rowcount := 0
	cellcount := 0
	columncount := 0
	brokencount := 0
	csvcount := 0

	for row := range e.ReadRows(e.Sheets[0]) {
		//fmt.Println("reading row")
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
			//for i := 0; i < len(record); i++ {
			//fmt.Println(i, record[i])
			//}
			/*if brokencount > 2 {
				break
			}*/
		} else {
			_ = csvwriter.Write(record)
			csvcount++
		}
	}
	fmt.Println("there were", brokencount, "broken records found out of", rowcount, "records !")
	fmt.Println("we wrote", csvcount, "csv records to", outfilepath)

	csvwriter.Flush()

	csvfile.Close()
	duration := time.Since(start)
	fmt.Println("duration:", duration)
	PrintMemUsage()

}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
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
