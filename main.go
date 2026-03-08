package main

import (
	"fmt"
)

func main() {
	staffMappingFilePath := "data/staff_mapping.csv"
	redemptionFilePath := "data/redemption.csv"

	InitializeCSVFiles(staffMappingFilePath, redemptionFilePath)

	var staffID string

	fmt.Print("Enter staff pass ID: ")
	fmt.Scanln(&staffID)
	redeemReward(staffMappingFilePath, redemptionFilePath, staffID)
}
