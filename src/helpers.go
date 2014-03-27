package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Read a whole file into the memory and store it as array of lines
func readLines(path string) (lines []string, err error) {
	var (
		file   *os.File
		part   []byte
		prefix bool
	)
	if file, err = os.Open(path); err != nil {
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		if len(part) > 0 && part[0] != '#' {
			buffer.Write(part)
			if !prefix {
				lines = append(lines, buffer.String())
				buffer.Reset()
			}
		}
	}
	if err == io.EOF {
		err = nil
	}
	return
}

func parseLine(line string) (item UnihanDataEntry, err error) {
	item = make([]string, 3)
	item = strings.Split(line, "\t")
	return item, nil
}

func parseUnihanFile(filename string, handlerName string, db *UnihanDB) {
	tx, err := db.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Commit()
	lines, err := readLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		item, err := parseLine(line)
		if err != nil {
			panic(err)
		}
		err = db.Insert(handlerName, tx, item)
		if err != nil {
			panic(err)
		}
	}
}

func hexToString(h string) (string, error) {
	return strconv.Unquote(`"\U` +
		fmt.Sprintf("%08s", h) + `"`)
}

type unihanValue struct {
	Value  string
	Remark string
}

/**
 * parse a Unihan database value field into a struct
 * represents the semantic values in it
 */
func parseUnihanValues(valueStr string) (returnVars []unihanValue, err error) {

	// split the value filed into values
	values := strings.Split(valueStr, " ")
	returnVars = make([]unihanValue, len(values), cap(values))

	// loop through values and parse
	for key, value := range values {
		fields := strings.Split(value, "<")
		if len(fields) == 0 {
			err = errors.New("Empty value")
			return
		} else if len(fields) == 1 {
			returnVars[key] = unihanValue{
				Value: fields[0],
			}
		} else {
			returnVars[key] = unihanValue{
				Value:  fields[0],
				Remark: fields[1],
			}
		}
	}
	return
}
