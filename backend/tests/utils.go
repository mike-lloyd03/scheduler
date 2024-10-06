package tests

import (
	"fmt"
	"testing"

	"github.com/pocketbase/pocketbase/models"
	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tokens"
)

type AuthHeaders struct {
	Org1Admin         map[string]string
	Org1Group1Admin   map[string]string
	Org1Member        map[string]string
	Org1Group1Manager map[string]string
	Org1Group1Member  map[string]string
}

func (a *AuthHeaders) Init(app *tests.TestApp) error {
	var err error

	a.Org1Admin, err = generateAuthHeader(app, "users", "org1Admin")
	if err != nil {
		return err
	}

	a.Org1Group1Admin, err = generateAuthHeader(app, "users", "org1group1Admin")
	if err != nil {
		return err
	}

	a.Org1Member, err = generateAuthHeader(app, "users", "org1Member")
	if err != nil {
		return err
	}

	a.Org1Group1Member, err = generateAuthHeader(app, "users", "org1group1Member")
	if err != nil {
		return err
	}

	return nil
}

type TestData struct {
	org1                  *models.Record
	org2                  *models.Record
	org1group1            *models.Record
	org2group1            *models.Record
	o1g1et1               *models.Record
	o1g2et1               *models.Record
	o2g1et1               *models.Record
	o1g1e1                *models.Record
	o1g2e1                *models.Record
	o2g1e1                *models.Record
	o1g1et1rt1            *models.Record
	o1g2et1rt1            *models.Record
	o2g1et1rt1            *models.Record
	onDeleteGroup         int
	onDeleteEventTemplate int
}

func (d *TestData) Init(app *tests.TestApp) error {
	var err error

	d.onDeleteGroup = 7
	d.onDeleteEventTemplate = 3

	d.org1, err = app.Dao().FindFirstRecordByFilter("orgs", "name='org1'")
	if err != nil {
		return err
	}

	d.org2, err = app.Dao().FindFirstRecordByFilter("orgs", "name='org2'")
	if err != nil {
		return err
	}

	d.org1group1, err = app.Dao().FindFirstRecordByFilter("groups", "name='org1group1'")
	if err != nil {
		return err
	}

	d.org2group1, err = app.Dao().FindFirstRecordByFilter("groups", "name='org2group1'")
	if err != nil {
		return err
	}

	d.o1g1et1, err = app.Dao().FindFirstRecordByFilter("event_templates", "name='o1g1et1'")
	if err != nil {
		return err
	}

	d.o1g2et1, err = app.Dao().FindFirstRecordByFilter("event_templates", "name='o1g2et1'")
	if err != nil {
		return err
	}

	d.o2g1et1, err = app.Dao().FindFirstRecordByFilter("event_templates", "name='o2g1et1'")
	if err != nil {
		return err
	}

	d.o1g1e1, err = app.Dao().FindFirstRecordByFilter("events", fmt.Sprintf("event_template='%s'", d.o1g1et1.Id))
	if err != nil {
		return err
	}

	d.o1g2e1, err = app.Dao().FindFirstRecordByFilter("events", fmt.Sprintf("event_template='%s'", d.o1g2et1.Id))
	if err != nil {
		return err
	}

	d.o2g1e1, err = app.Dao().FindFirstRecordByFilter("events", fmt.Sprintf("event_template='%s'", d.o2g1et1.Id))
	if err != nil {
		return err
	}

	d.o1g1et1rt1, err = app.Dao().FindFirstRecordByFilter("role_templates", "name='o1g1et1rt1'")
	if err != nil {
		return err
	}

	d.o1g2et1rt1, err = app.Dao().FindFirstRecordByFilter("role_templates", "name='o1g2et1rt1'")
	if err != nil {
		return err
	}

	d.o2g1et1rt1, err = app.Dao().FindFirstRecordByFilter("role_templates", "name='o2g1et1rt1'")
	if err != nil {
		return err
	}

	return nil
}

func generateTestApp(t *testing.T) *tests.TestApp {
	testApp, err := tests.NewTestApp(testDataDir)
	if err != nil {
		t.Fatal(err)
	}

	return testApp
}

func generateAdminToken(email string) (string, error) {
	app, err := tests.NewTestApp(testDataDir)
	if err != nil {
		return "", err
	}
	defer app.Cleanup()

	admin, err := app.Dao().FindAdminByEmail(email)
	if err != nil {
		return "", err
	}

	return tokens.NewAdminAuthToken(app, admin)
}

func generateAuthHeader(app *tests.TestApp, collectionNameOrId string, username string) (map[string]string, error) {
	header := map[string]string{}

	record, err := app.Dao().FindAuthRecordByUsername(collectionNameOrId, username)
	if err != nil {
		return header, err
	}

	token, err := tokens.NewRecordAuthToken(app, record)
	if err != nil {
		return header, err
	}

	header["Authorization"] = token

	return header, nil
}
