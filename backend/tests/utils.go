package tests

import (
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/pocketbase/pocketbase/tokens"
)

type AuthHeaders struct {
	Org1Admin        map[string]string
	Org1Group1Admin  map[string]string
	Org1Member       map[string]string
	Org1Group1Member map[string]string
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
