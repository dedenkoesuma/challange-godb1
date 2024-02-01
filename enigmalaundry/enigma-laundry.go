package enigmalaundry

import "time"

type Customer struct {
    Customer_id       int
	Customer_name string
	Phone_number  int
}
type Detail struct {
	Detail_id   int
	Date_entry  time.Time
	Date_finish time.Time
	Received_by string
}

type Sales struct{
	Sale_id int
	Services string
	Qty int
	Unit string
	Price string
	Amount string
}
