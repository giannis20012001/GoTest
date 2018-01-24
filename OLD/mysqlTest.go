package OLD

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

/**
 * Created by John Tsantilis (A.K.A lumi) on 14/12/2017.
 * @author John Tsantilis <i.tsantilis [at] yahoo [dot] com>
 */

func main() {
	fmt.Println("Go MySQL Tutorial")

	// Open up our database connection.
	db, err := sql.Open("mysql", "root:Floor53*@tcp(127.0.0.1:3306)/mysql")
	if err != nil {
		panic(err.Error())

	}

	defer db.Close()


	// Query the square-number of 13
	rows, err := db.Query("SHOW GLOBAL STATUS WHERE Variable_name = 'Bytes_sent'")
	if err != nil {
		panic(err.Error())

	}

	for rows.Next() {
		var yolo string
		var yolo2 int

		err = rows.Scan(&yolo, &yolo2)
		if err != nil {
			panic(err.Error())

		}

		fmt.Printf("Yolo: %d", yolo2)

	}

}