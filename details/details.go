package details

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

func AddNewDetail() {
	db := connectDb()
	defer db.Close()
	var err error

	var detail enigmalaundry.Detail
	detail.Detail_id = getUserInputInt("Enter Detail ID:")
	detail.Date_entry = getTimeInput("Enter Date Entry (e.g., '2006-01-02'):")
	detail.Date_finish = getTimeInput("Enter Date Finish (e.g., '2006-01-02'):")
	detail.Received_by = getUserInput("Enter Received By :")

	sqlStatement := "INSERT INTO details (detail_id, date_entry, date_finish, received_by) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(sqlStatement, detail.Detail_id, detail.Date_entry, detail.Date_finish, detail.Received_by)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully Insert Data!")
	}
}

func ViewDetail() []enigmalaundry.Detail {
	db := connectDb()
	defer db.Close()
	var err error

	var details []enigmalaundry.Detail

	rows, err := db.Query("SELECT detail_id, date_entry, date_finish, received_by FROM details")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var detail enigmalaundry.Detail
		err := rows.Scan(&detail.Detail_id, &detail.Date_entry, &detail.Date_finish, &detail.Received_by)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		details = append(details, detail)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil
	}

	return details
}

func UpdateDetail() {
	db := connectDb()
	defer db.Close()
	var err error

	DetailID := getUserInputInt("Enter Detail ID to update:")

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM details WHERE detail_id = $1", DetailID).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}

	if count == 0 {
		fmt.Println("Detail not found.")
		return
	}

	var detail enigmalaundry.Detail
	detail.Detail_id = DetailID // Set the Detail_id based on user input
	detail.Date_entry = getTimeInput("Enter new Date_entry:")
	detail.Date_finish = getTimeInput("Enter new Date_finish:")
	detail.Received_by = getUserInput("Enter new Received_by:")

	// Update detail in the database
	sqlStatement := "UPDATE details SET date_entry = $2, date_finish = $3, received_by = $4 WHERE detail_id = $1"
	_, err = db.Exec(sqlStatement, detail.Detail_id, detail.Date_entry, detail.Date_finish, detail.Received_by)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully updated detail information!")
	}
}

func DeleteDetail() {
	db := connectDb()
	defer db.Close()
	var err error

	DetailID := getUserInputInt("Enter Detail ID to update:")

	// Check if the customer exists
	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM details WHERE detail_id = $1", DetailID).Scan(&count)
	if err != nil {
		log.Fatal(err)
		return
	}

	if count == 0 {
		fmt.Println("Detail not found.")
		return
	}

	// Delete customer from the database
	sqlStatement := "DELETE FROM details WHERE detail_id = $1"
	_, err = db.Exec(sqlStatement, DetailID)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully deleted customer!")
	}
}