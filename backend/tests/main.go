package tests

import (
	"fmt"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tests"
)

const testDataDir = "../test_pb_data"
const OnRecordsListRequest = "OnRecordsListRequest"
const OnRecordViewRequest = "OnRecordViewRequest"
const OnModelAfterUpdate = "OnModelAfterUpdate"
const OnModelBeforeUpdate = "OnModelBeforeUpdate"
const OnRecordAfterUpdateRequest = "OnRecordAfterUpdateRequest"
const OnRecordBeforeUpdateRequest = "OnRecordBeforeUpdateRequest"

func clearTable(app *tests.TestApp, table string) error {
	records, err := app.Dao().FindRecordsByExpr(table)
	if err != nil {
		return err
	}
	for _, record := range records {
		err = app.Dao().DeleteRecord(record)
		if err != nil {
			return err
		}
	}
	return nil
}

func createRecord(app *tests.TestApp, collection *models.Collection, fields map[string]any) error {
	record := models.NewRecord(collection)
	record.Load(fields)
	fmt.Println("Creating record", record.Id)

	return app.Dao().SaveRecord(record)
}
