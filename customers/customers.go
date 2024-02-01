package customers

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"time"

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


func AddNewCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	var customer enigmalaundry.Customer
	customer.Customer_id = getUserInputInt("Enter Customer ID:")
	customer.Customer_name = getUserInput("Enter Customer Name:")
	customer.Phone_number = getUserInputInt("Enter Phone Number:")

	sqlStatement := "INSERT INTO customers (customer_id, customer_name, phone_number) VALUES ($1, $2, $3)"
	_, err = db.Exec(sqlStatement, customer.Customer_id, customer.Customer_name, customer.Phone_number)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func ViewCustomer() []enigmalaundry.Customer {
	db := connectDb()
	defer db.Close()
	var err error

	var customers []enigmalaundry.Customer

	rows, err := db.Query("SELECT customer_id,customer_name,phone_number FROM customers")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var customer enigmalaundry.Customer
		err := rows.Scan(&customer.Customer_id, &customer.Customer_name, &customer.Phone_number)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		customers = append(customers, customer)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return customers
}

func UpdateCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	// Get the customer ID to update
	customerID := getUserInputInt("Enter Customer ID to update:")

	// Check if the customer exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}

	if count == 0 {
		fmt.Println("Customer not found.")
		return
	}

	// Get updated customer information
	var customer enigmalaundry.Customer
	customer.Customer_name = getUserInput("Enter new Customer Name:")
	customer.Phone_number = getUserInputInt("Enter new Phone Number:")

	// Update customer in the database
	sqlStatement := "UPDATE customers SET customer_name = $2, phone_number = $3 WHERE customer_id = $1"
	_, err = db.Exec(sqlStatement, customer.Customer_name, customer.Phone_number, customerID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully updated customer information!")
	}
}
func DeleteCustomer() {
	db := connectDb()
	defer db.Close()
	var err error

	// Get the customer ID to delete
	customerID := getUserInputInt("Enter Customer ID to delete:")

	// Check if the customer exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM customers WHERE customer_id = $1", customerID).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}

	if count == 0 {
		fmt.Println("Customer not found.")
		return
	}

	// Delete customer from the database
	sqlStatement := "DELETE FROM customers WHERE customer_id = $1"
	_, err = db.Exec(sqlStatement, customerID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully deleted customer!")
	}
}