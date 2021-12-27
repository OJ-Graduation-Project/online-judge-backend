package main

import (
	"fmt"

	"github.com/OJ-Graduation-Project/online-judge-backend/internal/db"
)

func main() {
	defer db.Disconnect()
	db.InitializeDatabase()
	// db.InsertSubmission()
	sub, err := db.RetrieveSubmission(2)
	if err != nil {
		fmt.Print("error in retrieving")
	}
	fmt.Println(sub)
}
