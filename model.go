package main

import (
	"database/sql"
)

type phone struct {
	ID        int    `json:"id,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Company   string `json:"company,omitempty"`
	PhoneType string `json:"phone_type,omitempty"`
	UserID    string `json:"userId,omitempty"`
}

func (p *phone) createPhone(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO phones(phone, company, phoneType, userId) VALUES($1, $2, $3, $4) RETURNING id",
		p.Phone, p.Company, p.PhoneType, p.UserID).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func (p *phone) deletePhone(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM phones WHERE id=$1", p.ID)

	return err
}

func (p *phone) getPhone(db *sql.DB) error {
	return db.QueryRow("SELECT phone, company, phoneType, userId FROM phones WHERE id=$1",
		p.ID).Scan(&p.Phone, &p.Company, &p.PhoneType, &p.UserID)
}

func getPhones(db *sql.DB, start, count int) ([]phone, error) {
	rows, err := db.Query(
		"SELECT id, phone, company, phoneType, userId FROM phones LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	phones := []phone{}

	for rows.Next() {
		var p phone
		if err := rows.Scan(&p.ID, &p.Phone, &p.Company, &p.PhoneType, &p.UserID); err != nil {
			return nil, err
		}
		phones = append(phones, p)
	}

	return phones, nil
}

func (p *phone) updatePhone(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE phones SET phoneNumber=$1, company=$2, phoneType=$3, userId=$4 WHERE id=$5",
			p.Phone, p.Company, p.PhoneType, p.UserID, p.ID)

	return err
}