package main

import (
	"os"
	"testing"
)

func TestRedeemCases(t *testing.T) {
	staffFile := "test_staff.csv"
	redemptionFile := "test_redemption.csv"

	createSampleStaff(staffFile)
	defer os.Remove(staffFile)

	// case 1: staff not in csv
	file, _ := os.Create(redemptionFile)
	file.WriteString("team_name,redeemed_at\n")
	file.Close()
	defer os.Remove(redemptionFile)

	if redeemReward(staffFile, redemptionFile, "Sam") {
		t.Errorf("Expected false for unknown staff")
	}

	// case 2: staff in csv, team has already redeemed
	file, _ = os.Create(redemptionFile)
	file.WriteString("team_name,redeemed_at\nTeamA,1700000000000\n") // pre-existing redemption
	file.Close()

	if redeemReward(staffFile, redemptionFile, "Staff1") {
		t.Errorf("Expected TeamA to have already redeemed")
	}

	// case 3: staff in csv, team not redeemed
	file, _ = os.Create(redemptionFile)
	file.WriteString("team_name,redeemed_at\n") // empty redemption
	file.Close()

	if !redeemReward(staffFile, redemptionFile, "Staff1") {
		t.Errorf("Expected TeamA to be added to redemption CSV")
	}

	os.Remove(redemptionFile)
}
