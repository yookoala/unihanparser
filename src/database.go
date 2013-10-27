package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type UnihanDB struct {
	Filename string
	DB       *sql.DB
	conn     bool
	handlers map[string]UnihanFileHandler
}

func (uni *UnihanDB) InitDb() (err error) {
	return
}

func (uni *UnihanDB) Register(name string, h UnihanFileHandler) (err error) {
	if uni.conn != true {
		err := uni.Open()
		if err != nil {
			return err
		}
		uni.handlers = make(map[string]UnihanFileHandler)
	}

	uni.handlers[name] = h

	return h.Init(uni)
}

func (uni *UnihanDB) Insert(name string, tx *sql.Tx, item []string) (err error) {
	if handler, ok := uni.handlers[name]; ok {
		handler.Insert(tx, item)
		return nil
	}
	return errors.New(fmt.Sprintf("Handler \"%s\" not found", name))
}

func (uni *UnihanDB) Open() (err error) {
	db, err := sql.Open("sqlite3", uni.Filename)
	if err != nil {
		return
	}
	uni.DB = db
	uni.conn = true
	return nil
}

func (uni *UnihanDB) Close() (err error) {
	uni.DB.Close()
	return nil
}
