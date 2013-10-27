package main

import (
	"bufio"
	"bytes"
	"io"
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

func parseLines(lines []string) (items [][]string) {
	items = make([][]string, len(lines))
	for i, line := range lines {
		items[i] = make([]string, 3)
		items[i] = strings.Split(line, "\t")
	}
	return
}

func main() {
	lines, err := readLines("./Unihan-6.1/Unihan_IRGSources.txt")
	if err != nil {
		panic(err)
	}
	items := parseLines(lines)

	print(items[0][0])
}
