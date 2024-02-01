package sales

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"git.enigmacamp.com/enigma-20/deden/challange-godb1/enigmalaundry"
)


const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "Golang-Database"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func connectDb() *sql.DB {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to database")
	}
	return db
}

func getUserInput(prompt string) string {
	fmt.Print(prompt + " ")
	var input string
	fmt.Scanln(&input)
	return input
}
func getUserInputInt(prompt string) int {
	inputString := getUserInput(prompt)
	num, err := strconv.Atoi(inputString)
	if err != nil {
		log.Fatal(err)
	}
	return num
}


func AddNewSale() {
	db := connectDb()
	defer db.Close()
	var err error

	var sales enigmalaundry.Sales
	sales.Sale_id = getUserInputInt("Enter Sales ID:")
	sales.Services = getUserInput("Enter Services:")
	sales.Qty = getUserInputInt("Enter Quantity:")
	sales.Unit = getUserInput("Enter Unit:")
	sales.Price = getUserInput("Enter Price:")
	sales.Amount = getUserInput("Enter Amount:")

	sqlStatement := "INSERT INTO sales (sale_id, services, qty, unit, price, amount) VALUES ($1, $2, $3, $4, $5, $6)"
	_, err = db.Exec(sqlStatement, sales.Sale_id,sales.Services,sales.Qty,sales.Unit,sales.Price,sales.Amount)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func ViewSale() []enigmalaundry.Sales {
	db := connectDb()
	defer db.Close()
	var err error

	var Sales []enigmalaundry.Sales

	rows, err := db.Query("SELECT sale_id, services, qty, unit, price, amount FROM sales")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var sale enigmalaundry.Sales
		err := rows.Scan(&sale.Sale_id, &sale.Services, &sale.Qty, &sale.Unit, &sale.Price, &sale.Amount)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		Sales = append(Sales, sale)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return Sales
}
