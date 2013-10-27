package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
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

func main() {
	dir := "./data/Unihan-6.1"
	db := UnihanDB{
		Filename: "./data/unihan.db",
	}
	db.Register("Variants", VariantsHandler{})
	parseUnihanFile(dir+"/Unihan_Variants.txt", "Variants", &db)
}