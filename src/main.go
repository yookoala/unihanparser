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

	db.Register("DictionaryIndices", GenericDataHandler{TableName: "DictionaryIndices"})
	db.Register("DictionaryLikeData", GenericDataHandler{TableName: "DictionaryLikeData"})
	db.Register("IRGSources", GenericDataHandler{TableName: "IRGSources"})
	db.Register("NumericValues", GenericDataHandler{TableName: "NumericValues"})
	db.Register("OtherMappings", GenericDataHandler{TableName: "OtherMappings"})
	db.Register("RadicalStrokeCounts", RadicalStrokeCountsHandler{})
	db.Register("Readings", GenericDataHandler{TableName: "Readings"})
	db.Register("Variants", VariantsHandler{})

	parseUnihanFile(dir+"/Unihan_DictionaryIndices.txt", "DictionaryIndices", &db)
	parseUnihanFile(dir+"/Unihan_DictionaryLikeData.txt", "DictionaryLikeData", &db)
	parseUnihanFile(dir+"/Unihan_IRGSources.txt", "IRGSources", &db)
	parseUnihanFile(dir+"/Unihan_NumericValues.txt", "NumericValues", &db)
	parseUnihanFile(dir+"/Unihan_OtherMappings.txt", "OtherMappings", &db)
	parseUnihanFile(dir+"/Unihan_RadicalStrokeCounts.txt", "RadicalStrokeCounts", &db)
	parseUnihanFile(dir+"/Unihan_Readings.txt", "Readings", &db)
	parseUnihanFile(dir+"/Unihan_Variants.txt", "Variants", &db)
}
