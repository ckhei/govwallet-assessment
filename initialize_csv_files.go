package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func InitializeCSVFiles(staffMappingFilePath string, redemptionFilePath string) {
	InitializeStaffCSV(staffMappingFilePath)
	InitializeRedemptionCSV(redemptionFilePath)
}

func InitializeRedemptionCSV(filePath string) {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.Size() == 0 {
		data := [][]string{
			{"team_name", "redeemed_at"},
		}
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating redemption CSV:", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.WriteAll(data)
	}
}

func InitializeStaffCSV(filePath string) {
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) || info.Size() == 0 {
		createSampleStaff(filePath)
	}
}

func createSampleStaff(filePath string) {

	data := [][]string{
		{"staff_pass_id", "team_name", "created_at"},
		{"Staff1", "TeamA", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"Staff2", "TeamB", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"Staff3", "TeamC", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"Staff4", "TeamD", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"Staff5", "TeamE", fmt.Sprintf("%d", time.Now().UnixMilli())},
		{"Staff6", "TeamF", fmt.Sprintf("%d", time.Now().UnixMilli())},
	}

	file, _ := os.Create(filePath)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.WriteAll(data)
}
