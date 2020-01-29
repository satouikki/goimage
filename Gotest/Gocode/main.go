package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Info struct {
	ID   int
	Host string
	Date string
}

const DRIVER = "mysql"
const DSN = "golang-test-user:golang-test-pass@tcp(mysql-container:3306)/golang-test-database"

func main() {

	http.HandleFunc("/app", gotest)
	http.ListenAndServe(":8080", nil)
}

func gotest(w http.ResponseWriter, r *http.Request) {

	db, err := sql.Open(DRIVER, DSN)
	if err != nil {
		fmt.Println("Openエラー")
	} else {
		fmt.Println("OpenOK！")
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("接続失敗！")
	} else {
		fmt.Println("接続OK！")
	}
	hos, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	times := time.Now()
	tim := times.String()
	fmt.Println(hos)
	fmt.Println(tim)

	var info Info
	err = db.QueryRow("select * from data WHERE host=?", hos).Scan(&info.ID, &info.Host, &info.Date)
	switch {
	case err == sql.ErrNoRows:
		fmt.Println("No Record,Creating New Record")
		ins, err := db.Prepare("INSERT INTO data(host,date) VALUES(?,?)")
		if err != nil {
			log.Fatal(err)
		}
		ins.Exec(hos, tim)
		fmt.Fprintf(w, "Hello, %s,This is Your First Login, Time is %#v", hos, tim)
	case err != nil:
		panic(err.Error())
	default:
		fmt.Println(info.ID, info.Host, info.Date)
		upd, err := db.Prepare("UPDATE data SET date = ? WHERE host = ? ")
		if err != nil {
			log.Fatal(err)
		}
		upd.Exec(times, hos)
		fmt.Fprintf(w, "Hello, %s,Your Latest Login is %#v", hos, info.Date)

	}

	db.Close()
}
