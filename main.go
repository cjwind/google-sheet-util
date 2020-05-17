package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"
	"log"
)

type GoogleSheet struct {
	srv           *sheets.Service
	spreadsheetId string
}

func (gs *GoogleSheet) Init(spreadsheetId string) {
	gs.spreadsheetId = spreadsheetId

	var err error
	ctx := context.Background()
	gs.srv, err = sheets.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to new sheets service: %v", err)
	}
}

func main() {
	spreadsheetId := "12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0"

	var gs GoogleSheet
	gs.Init(spreadsheetId)

	// https://docs.google.com/spreadsheets/d/12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0/edit
	readRange := "Table!A2:E"

	values := gs.Read(readRange)
	fmt.Println(values)

	var writeValues [][]interface{}
	row := []interface{}{"AAA", "BBB"}
	writeValues = append(writeValues, row)
	gs.Write("Table", writeValues)

	var updateValues [][]interface{}
	row = []interface{}{"BBB", "CCC", "DDD"}
	updateValues = append(updateValues, row)
	gs.Update("Table!A3", updateValues)

	clearRange := "Table!A3:E"
	gs.Clear(clearRange)
}

func (gs *GoogleSheet) Clear(clearRange string) {
	// rb has type *ClearValuesRequest
	rb := &sheets.ClearValuesRequest{}

	_, err := gs.srv.Spreadsheets.Values.Clear(gs.spreadsheetId, clearRange, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) Update(updateRange string, updateValues [][]interface{}) {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         updateValues,
	}

	_, err := gs.srv.Spreadsheets.Values.Update(gs.spreadsheetId, updateRange, rb).ValueInputOption(valueInputOption).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) Write(writeRange string, values [][]interface{}) {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         values,
	}

	_, err := gs.srv.Spreadsheets.Values.Append(gs.spreadsheetId, writeRange, rb).ValueInputOption(valueInputOption).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) Read(readRange string) [][]interface{} {
	resp, err := gs.srv.Spreadsheets.Values.Get(gs.spreadsheetId, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Values
}
