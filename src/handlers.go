package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type UnihanDataEntry []string

type UnihanFileHandler interface {
	Init(uni *UnihanDB) (err error)
	ParseLine(line string) (item UnihanDataEntry, err error)
	Insert(tx *sql.Tx, item UnihanDataEntry) (err error)
}

// hander for the Unihan_Variants.txt
type VariantsHandler struct {
}

func (h VariantsHandler) Init(uni *UnihanDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS Variants (
			id INTEGER PRIMARY KEY,
			unicode TEXT,
			character TEXT,
			variant_type TEXT,
			variant_unicode TEXT,
			variant_character TEXT,
			variant_remark TEXT
		)
	`
	_, err = uni.DB.Exec(query)

	return
}

func (h VariantsHandler) ParseLine(line string) (item UnihanDataEntry, err error) {
	return parseLine(line)
}

func (h VariantsHandler) Insert(tx *sql.Tx, item UnihanDataEntry) (err error) {
	stmt, err := tx.Prepare(`INSERT INTO Variants (
		unicode,
		character,
		variant_type,
		variant_unicode,
		variant_character,
		variant_remark
	) VALUES (
		?, ?, ?, ?, ?, ?
	)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(item[0], item[0], item[1], item[2], item[2], item[2])
	return
}
