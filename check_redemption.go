package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func redeemReward(staffMappingFilePath string, redemptionFilePath string, staffID string) {
	teamName := checkStaff(staffMappingFilePath, staffID)
	hasRedeemed := checkRedemption(redemptionFilePath, teamName)

	if !hasRedeemed {
		file, err := os.OpenFile(redemptionFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Printf("Error opening redemption CSV for writing: %v\n", err)
			return
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		defer writer.Flush()

		writer.Write([]string{teamName, fmt.Sprintf("%d", time.Now().UnixMilli())})
		fmt.Printf("Team %s has successfully redeemed the reward.\n", teamName)
	} else {
		fmt.Printf("Team %s is unable to redeem the reward.\n", teamName)
	}
}

func checkStaff(filePath string, staffID string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening staff CSV: %v\n", err)
		return ""
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading staff CSV: %v\n", err)
		return ""
	}

	for i, record := range records {
		if i == 0 { // header row
			continue
		}

		if record[0] == staffID {
			teamName := record[1]
			return teamName
		}
	}
	return ""
}

func checkRedemption(filePath string, teamName string) bool {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening redemption CSV: %v\n", err)
		return true
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading redemption CSV: %v\n", err)
		return true
	}

	for i, record := range records {
		if i == 0 { // header row
			continue
		}

		if record[0] == teamName {
			return true
		}
	}
	return false
}
