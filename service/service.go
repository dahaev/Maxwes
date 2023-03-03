package service

import (
	"encoding/csv"
	"log"
	"maxwes_group/models"
	"os"
	"strconv"
)

// nice file name

const (
	fileName = "ueba.csv"
)

// Openening csv file...

func ReadCsvFile() [][]string {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal("File error")
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()

	if err != nil {
		log.Fatal("File error")
	}

	return records[1:][:]
}

// Find member by id... and returnig model.Membership...

func FindMembership(id int) (*models.Membership, error) {

	record := ReadCsvFile()

	var idx_number int

	for idx, value := range record {
		number, err := strconv.Atoi(value[1])
		if err != nil {
			log.Fatal("Value is not integer")
			return nil, err
		}

		if number == id {
			idx_number = idx
			break
		}
	}

	// Kostil ebaniy :(

	var member models.Membership

	member.Number = record[idx_number][0]
	member.ID = record[idx_number][1]
	member.UID = record[idx_number][2]
	member.Domain = record[idx_number][3]
	member.CN = record[idx_number][4]
	member.Department = record[idx_number][5]
	member.Title = record[idx_number][6]
	member.Who = record[idx_number][7]
	member.Logon_count = record[idx_number][8]
	member.Num_logons7 = record[idx_number][9]
	member.Num_share7 = record[idx_number][10]
	member.Num_file7 = record[idx_number][11]
	member.Num_ad7 = record[idx_number][12]
	member.Num_n7 = record[idx_number][13]
	member.Num_logons14 = record[idx_number][14]
	member.Num_share14 = record[idx_number][15]
	member.Num_file14 = record[idx_number][16]
	member.Num_ad14 = record[idx_number][17]
	member.Num_n14 = record[idx_number][18]
	member.Num_logons30 = record[idx_number][19]
	member.Num_share30 = record[idx_number][20]
	member.Num_file30 = record[idx_number][21]
	member.Num_ad30 = record[idx_number][22]
	member.Num_n30 = record[idx_number][23]
	member.Num_logons150 = record[idx_number][24]
	member.Num_share150 = record[idx_number][25]
	member.Num_file150 = record[idx_number][26]
	member.Num_ad150 = record[idx_number][27]
	member.Num_n150 = record[idx_number][28]
	member.Num_logons365 = record[idx_number][29]
	member.Num_share365 = record[idx_number][30]
	member.Num_file365 = record[idx_number][31]
	member.Num_ad365 = record[idx_number][32]
	member.Num_n365 = record[idx_number][33]
	member.Has_user_principal_name = record[idx_number][34]
	member.Has_mail = record[idx_number][35]
	member.Has_phone = record[idx_number][36]
	member.Flag_disabled = record[idx_number][37]
	member.Flag_lockout = record[idx_number][38]
	member.Flag_password_not_required = record[idx_number][39]
	member.Flag_password_cant_change = record[idx_number][40]
	member.Flag_dont_expire_password = record[idx_number][41]
	member.Owned_files = record[idx_number][42]
	member.Num_mailboxes = record[idx_number][43]
	member.Num_member_of_groups = record[idx_number][44]
	member.Num_member_of_indirect_groups = record[idx_number][45]
	member.Member_of_indirect_groups_ids = record[idx_number][46]
	member.Member_of_groups_ids = record[idx_number][47]
	member.Is_admin = record[idx_number][48]
	member.Is_service = record[idx_number][49]

	return &member, nil
}

func MembershipRequest(members_id []string) []models.Membership {

	var members []models.Membership

	for _, id := range members_id {

		id, _ := strconv.Atoi(id)
		member, err := FindMembership(id)

		if err != nil {
			continue
		}
		if member == nil {
			continue
		}
		members = append(members, *member)
	}

	return members
}
