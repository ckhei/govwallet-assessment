package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func InitializeStaffCSV(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createSampleStaff(filePath)
	}
}

func createSampleStaff(filePath string) {

	data := [][]string{
		{"staff_pass_id", "team_name", "created_at"},
		{"A123", "TeamAlpha", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"B456", "TeamBeta", fmt.Sprintf("%d", time.Now().UnixMilli())},
	}

	file, _ := os.Create(filePath)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(data)
}
