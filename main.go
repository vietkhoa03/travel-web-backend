package main

import (
	"fmt"
	connectDB "travel-web-backend/database"
)

func main() {
	travelDB := connectDB.ConnectTravelDB()

	// Ví dụ: in ra tên DB để kiểm tra
	fmt.Println("Database name:", travelDB.Name())
}