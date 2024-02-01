package main

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

	"git.enigmacamp.com/enigma-20/deden/challange-godb1/customers"
	"git.enigmacamp.com/enigma-20/deden/challange-godb1/details"
	"git.enigmacamp.com/enigma-20/deden/challange-godb1/sales"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "Golang-Database"
)

var psqlInfo = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

func main() {
	MainMenu()
}

func MainMenu(){
	fmt.Println("1. Menu Customer")
	fmt.Println("2. Menu Detail")
	fmt.Println("3. Menu Sale")

	choice := getUserInputInt("Enter your choice (1-3):")

	switch choice {
	case 1:
		MenuofCustomer()
	case 2:
		MenuOfDetails()
	case 3:
		MenuOfSales()
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
	} 
}

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
func getTimeInput(prompt string) time.Time {
	var input string

	for {
		fmt.Print(prompt + " (format: YYYY-MM-DD): ")
		fmt.Scanln(&input)

		// Melakukan parsing string ke tipe data time.Time
		dateEntry, err := time.Parse("2006-01-02", input)
		if err == nil {
			return dateEntry
		}
		fmt.Println("Format tanggal tidak valid. Silakan coba lagi.")
	}
}

func MenuOfSales(){
	fmt.Println("1. View Sale")
	fmt.Println("2. Add New Sale")
	fmt.Println("0. Back")

	choice := getUserInputInt("Enter your choice (1-2):")

	switch choice {
	case 0:
		MainMenu()
	case 1:
		sales := sales.ViewSale()
		for _, sale := range sales {
			fmt.Println(sale.Sale_id,sale.Services,sale.Qty,sale.Unit,sale.Price,sale.Amount)
		}
	case 2:
		sales.AddNewSale()
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
	} 
}

func MenuOfDetails(){
	fmt.Println("1. View Detail")
	fmt.Println("2. Add New Detail")
	fmt.Println("3. Update Detail")
	fmt.Println("4. Delete Detail")
	fmt.Println("0. Back")

	choice := getUserInputInt("Enter your choice (1-4):")

	switch choice {
	case 0:
		MainMenu()
	case 1:
		details := details.ViewDetail()
		for _, detail := range details {
			fmt.Println(detail.Detail_id, detail.Date_entry, detail.Date_finish, detail.Received_by)
		}
	case 2:
		details.AddNewDetail()
	case 3:
		details.UpdateDetail()
	case 4:
		details.DeleteDetail()
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
	} 
}

func MenuofCustomer(){
	fmt.Println("1. View Customer")
	fmt.Println("2. Add New Customer")
	fmt.Println("3. Update Customer")
	fmt.Println("4. Delete Customer")
	fmt.Println("0. Back")

	choice := getUserInputInt("Enter your choice (1-4):")

	switch choice {
	case 0:
		MainMenu()
	case 1:
		customers := customers.ViewCustomer()
		for _, customer := range customers {
			fmt.Println(customer.Customer_id, customer.Customer_name, customer.Phone_number)
		}
	case 2:
		customers.AddNewCustomer()
	case 3:
		customers.UpdateCustomer()
	case 4:
		customers.DeleteCustomer()
	default:
		fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
	} 
}