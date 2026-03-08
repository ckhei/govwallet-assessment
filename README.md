# govwallet-assessment
## Assumptions:
- Each staff ID is unique and belongs to exactly one team.
- Redemption data is stored in a CSV file, and timestamps are in epoch milliseconds.
- Staff ID input is required from the user to redeem a reward.
- Unit tests use temporary CSVs to avoid changing actual data.
- Unit tests cover the 3 key scenarios for `redeemReward()`:
    1. Staff not in staff mapping
    2. Team has already redeemed before
    3. Team has not redeem before

## How to run:
1. Clone the repository: git clone https://github.com/ckhei/govwallet-assessment.git
2. cd govwallet-assessment
3. go run .
4. Enter StaffID
4. go test -v (to run unit tests)