package main

import (
	"golang.org/x/net/context"
	"google.golang.org/api/sheets/v4"
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

func (gs *GoogleSheet) Write(writeRange string, values [][]interface{}) error {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         values,
	}

	_, err := gs.srv.Spreadsheets.Values.Append(gs.spreadsheetId, writeRange, rb).ValueInputOption(valueInputOption).Do()

	return err
}

func (gs *GoogleSheet) Update(updateRange string, updateValues [][]interface{}) error {
	valueInputOption := "RAW"
	rb := &sheets.ValueRange{
		MajorDimension: "ROWS",
		Values:         updateValues,
	}

	_, err := gs.srv.Spreadsheets.Values.Update(gs.spreadsheetId, updateRange, rb).ValueInputOption(valueInputOption).Do()

	return err
}

func (gs *GoogleSheet) Clear(clearRange string) error {
	// rb has type *ClearValuesRequest
	rb := &sheets.ClearValuesRequest{}

	_, err := gs.srv.Spreadsheets.Values.Clear(gs.spreadsheetId, clearRange, rb).Do()

	return err
}
