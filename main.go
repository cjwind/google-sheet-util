package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"
	"log"
)

type GoogleSheet struct {
	srv *sheets.Service
}

func (gs *GoogleSheet) Init() {
	var err error

	ctx := context.Background()
	gs.srv, err = sheets.NewService(ctx)
	if err != nil {
		log.Fatalf("Unable to new sheets service: %v", err)
	}
}

func main() {
	var gs GoogleSheet
	gs.Init()

	// https://docs.google.com/spreadsheets/d/12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0/edit
	spreadsheetId := "12Luq-VG23UxdcIhfmNCtW_BL4fpHPUwUK-cfwRJyJx0"
	readRange := "Table!A2:E"

	values := gs.Read(spreadsheetId, readRange)
	fmt.Println(values)

	var writeValues [][]interface{}
	row := []interface{}{"AAA", "BBB"}
	writeValues = append(writeValues, row)
	gs.write(spreadsheetId, "Table", writeValues)

	var updateValues [][]interface{}
	row = []interface{}{"BBB", "CCC", "DDD"}
	updateValues = append(updateValues, row)
	gs.update(spreadsheetId, "Table!A3", updateValues)

	clearRange := "Table!A3:E"
	gs.clear(spreadsheetId, clearRange)
}

func (gs *GoogleSheet) clear(spreadsheetId string, clearRange string) {
	// rb has type *ClearValuesRequest
	rb := &sheets.ClearValuesRequest{}

	_, err := gs.srv.Spreadsheets.Values.Clear(spreadsheetId, clearRange, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) update(spreadsheetId string, updateRange string, updateValues [][]interface{}) {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         updateValues,
	}

	_, err := gs.srv.Spreadsheets.Values.Update(spreadsheetId, updateRange, rb).ValueInputOption(valueInputOption).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) write(spreadsheetId string, writeRange string, values [][]interface{}) {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         values,
	}

	_, err := gs.srv.Spreadsheets.Values.Append(spreadsheetId, writeRange, rb).ValueInputOption(valueInputOption).Do()
	if err != nil {
		log.Fatal(err)
	}
}

func (gs *GoogleSheet) Read(spreadsheetId string, readRange string) [][]interface{} {
	resp, err := gs.srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()

	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	return resp.Values
}
