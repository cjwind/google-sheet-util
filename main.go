package main

import (
	"fmt"
	"log"
)

func main() {
	// https://docs.google.com/spreadsheets/d/12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0/edit
	spreadsheetId := "12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0"

	var gs GoogleSheet
	if err := gs.Init(spreadsheetId); err != nil {
		log.Fatalf("Unable to init google sheet api: %v\n Run source env?", err)
	}

	readRange := "Table!A2:E"
	values, err := gs.Read(readRange)
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	fmt.Println(values)

	var writeValues [][]interface{}
	row := []interface{}{"AAA", "BBB"}
	writeValues = append(writeValues, row)
	err = gs.Write("Table", writeValues)
	if err != nil {
		log.Fatal(err)
	}

	var updateValues [][]interface{}
	row = []interface{}{"BBB", "CCC", "DDD"}
	updateValues = append(updateValues, row)
	err = gs.Update("Table!A3", updateValues)
	if err != nil {
		log.Fatal(err)
	}

	clearRange := "Table!A3:E"
	err = gs.Clear(clearRange)
	if err != nil {
		log.Fatal(err)
	}
}
