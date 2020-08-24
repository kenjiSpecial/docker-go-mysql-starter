package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "helloworld\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

const (
	DRIVER_NAME = "mysql" // ドライバ名(mysql固定)
	// user:password@tcp(container-name:port)/dbname ※mysql はデフォルトで用意されているDB
	DATA_SOURCE_NAME = "root:password@tcp(mysql:3306)/database"
)

func main() {
	db, err := sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	} else {
		fmt.Println("Database created sucessfully")
	}
	defer db.Close()

	// name := "database"

	_, err = db.Exec("CREATE TABLE example ( id integer, data varchar(32) )")
	if err != nil {
		panic(err)
	}

}
