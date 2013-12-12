package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strings"
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

	// turn a given unicode into character
	character, err := hexToString(item[0][2:])
	if err != nil {
		log.Fatal(err)
	}

	// read the unihan values
	item[2] = strings.Trim(item[2], " ")
	unihanVars, err := parseUnihanValues(item[2])
	if err != nil {
		log.Fatal(err)
	}

	// loop through all values
	for _, unihanVar := range unihanVars {
		// turn a given unicode into variant_character
		variant_character, err := hexToString(unihanVar.Value[2:])
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(item[0], character, item[1], unihanVar.Value, variant_character, unihanVar.Remark)
	}
	return
}

// hander for the Unihan_RadicalStrokeCounts.txt
type RadicalStrokeCountsHandler struct {
}

func (h RadicalStrokeCountsHandler) Init(uni *UnihanDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS RadicalStrokeCounts (
			id INTEGER PRIMARY KEY,
			unicode TEXT,
			character TEXT,
			type TEXT,
			data TEXT
		)
	`
	_, err = uni.DB.Exec(query)

	return
}

func (h RadicalStrokeCountsHandler) ParseLine(line string) (item UnihanDataEntry, err error) {
	return parseLine(line)
}

func (h RadicalStrokeCountsHandler) Insert(tx *sql.Tx, item UnihanDataEntry) (err error) {
	stmt, err := tx.Prepare(`INSERT INTO RadicalStrokeCounts (
		unicode,
		character,
		type,
		data
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// turn a given unicode into character
	character, err := hexToString(item[0][2:])
	if err != nil {
		log.Fatal(err)
	}

	// read the unihan values
	item[2] = strings.Trim(item[2], " ")
	unihanVars, err := parseUnihanValues(item[2])
	if err != nil {
		log.Fatal(err)
	}

	// loop through all values
	for _, unihanVar := range unihanVars {
		_, err = stmt.Exec(item[0], character, item[1], unihanVar.Value)
	}
	return
}

// hander for the generic Unihan data files
type GenericDataHandler struct {
	TableName string
}

func (h GenericDataHandler) Init(uni *UnihanDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS ` + h.TableName + ` (
			id INTEGER PRIMARY KEY,
			unicode TEXT,
			character TEXT,
			type TEXT,
			data TEXT
		)
	`
	_, err = uni.DB.Exec(query)

	return
}

func (h GenericDataHandler) ParseLine(line string) (item UnihanDataEntry, err error) {
	return parseLine(line)
}

func (h GenericDataHandler) Insert(tx *sql.Tx, item UnihanDataEntry) (err error) {
	stmt, err := tx.Prepare(`INSERT INTO ` + h.TableName + ` (
		unicode,
		character,
		type,
		data
	) VALUES (
		?, ?, ?, ?
	)`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// turn a given unicode into character
	character, err := hexToString(item[0][2:])
	if err != nil {
		log.Fatal(err)
	}

	// read the unihan values
	item[2] = strings.Trim(item[2], " ")
	unihanVars, err := parseUnihanValues(item[2])
	if err != nil {
		log.Fatal(err)
	}

	// loop through all values
	for _, unihanVar := range unihanVars {
		// turn a given unicode into variant_character
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmt.Exec(item[0], character, item[1], unihanVar.Value)
	}
	return
}
