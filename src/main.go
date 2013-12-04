package main

import ()

func main() {
	dir := "./data/Unihan-6.3"
	db := UnihanDB{
		Filename: "./data/unihan.db",
	}

	db.Register("Variants", VariantsHandler{})
	db.Register("DictionaryLikeData", DictionaryLikeDataHandler{})
	db.Register("OtherMappings", OtherMappingsHandler{})

	parseUnihanFile(dir+"/Unihan_Variants.txt", "Variants", &db)
	parseUnihanFile(dir+"/Unihan_DictionaryLikeData.txt", "DictionaryLikeData", &db)
	parseUnihanFile(dir+"/Unihan_OtherMappings.txt", "OtherMappings", &db)
}
