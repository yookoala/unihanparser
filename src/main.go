package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var dir, dbfn string

	flag.StringVar(&dir, "f", "./data/Unihan", "Folder containing Unihan database files")
	flag.StringVar(&dbfn, "d", "./data/unihan.db", "Output database filename")
	flag.Parse()

	// check if Unihan database folder is a folder
	finfo, err := os.Stat(dir)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	} else if !finfo.IsDir() {
		fmt.Println("The directory \"" + dir + "\" should be a directory")
		os.Exit(1)
	}

	db := UnihanDB{
		Filename: dbfn,
	}

	db.Register("Variants", VariantsHandler{})
	db.Register("DictionaryLikeData", GenericDataHandler{TableName: "DictionaryLikeData"})
	db.Register("OtherMappings", GenericDataHandler{TableName: "OtherMappings"})
	db.Register("IRGSources", GenericDataHandler{TableName: "IRGSources"})

	parseUnihanFile(dir+"/Unihan_Variants.txt", "Variants", &db)
	parseUnihanFile(dir+"/Unihan_DictionaryLikeData.txt", "DictionaryLikeData", &db)
	parseUnihanFile(dir+"/Unihan_OtherMappings.txt", "OtherMappings", &db)
	parseUnihanFile(dir+"/Unihan_IRGSources.txt", "IRGSources", &db)
}
