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

// hander for the Unihan_DictionaryLikeData.txt
type DictionaryLikeDataHandler struct {
}

func (h DictionaryLikeDataHandler) Init(uni *UnihanDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS DictionaryLikeData (
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

func (h DictionaryLikeDataHandler) ParseLine(line string) (item UnihanDataEntry, err error) {
	return parseLine(line)
}

func (h DictionaryLikeDataHandler) Insert(tx *sql.Tx, item UnihanDataEntry) (err error) {
	stmt, err := tx.Prepare(`INSERT INTO DictionaryLikeData (
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

// hander for the Unihan_OtherMappings.txt
type OtherMappingsHandler struct {
}

func (h OtherMappingsHandler) Init(uni *UnihanDB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS OtherMappings (
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

func (h OtherMappingsHandler) ParseLine(line string) (item UnihanDataEntry, err error) {
	return parseLine(line)
}

func (h OtherMappingsHandler) Insert(tx *sql.Tx, item UnihanDataEntry) (err error) {
	stmt, err := tx.Prepare(`INSERT INTO OtherMappings (
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
