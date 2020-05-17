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

func (gs *GoogleSheet) Init(spreadsheetId string) (err error) {
	gs.spreadsheetId = spreadsheetId

	ctx := context.Background()
	gs.srv, err = sheets.NewService(ctx)

	return err
}

func (gs *GoogleSheet) Read(readRange string) ([][]interface{}, error) {
	resp, err := gs.srv.Spreadsheets.Values.Get(gs.spreadsheetId, readRange).Do()

	if err != nil {
		return nil, err
	}

	return resp.Values, err
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

func (gs *GoogleSheet) Clear(clearRange string) {
	// rb has type *ClearValuesRequest
	rb := &sheets.ClearValuesRequest{}

	_, err := gs.srv.Spreadsheets.Values.Clear(gs.spreadsheetId, clearRange, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// https://docs.google.com/spreadsheets/d/12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0/edit
	spreadsheetId := "12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0"

	var gs GoogleSheet
	err := gs.Init(spreadsheetId)
	if err != nil {
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
	gs.Write("Table", writeValues)

	var updateValues [][]interface{}
	row = []interface{}{"BBB", "CCC", "DDD"}
	updateValues = append(updateValues, row)
	gs.Update("Table!A3", updateValues)

	clearRange := "Table!A3:E"
	gs.Clear(clearRange)
}
