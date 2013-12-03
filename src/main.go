package main

import ()

func main() {
	dir := "./data/Unihan-6.3"
	db := UnihanDB{
		Filename: "./data/unihan.db",
	}
	db.Register("Variants", VariantsHandler{})
	parseUnihanFile(dir+"/Unihan_Variants.txt", "Variants", &db)
}
