package tests

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
)

func TestGroups(t *testing.T) {
	app := generateTestApp(t)

	org1AdminHeader, err := generateAuthHeader(app, "users", "org1Admin")
	if err != nil {
		t.Fatal(err)
	}

	org1MemberHeader, err := generateAuthHeader(app, "users", "org1Member")
	if err != nil {
		t.Fatal(err)
	}

	org1, err := app.Dao().FindRecordsByFilter("orgs", "name='org1'", "", 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	org2, err := app.Dao().FindRecordsByFilter("orgs", "name='org2'", "", 1, 0)
	if err != nil {
		t.Fatal(err)
	}

	scenarios := []tests.ApiScenario{
		// -----------
		// List
		// -----------
		{
			Name:            "unauth cannot list orgs",
			Method:          http.MethodGet,
			Url:             "/api/collections/orgs/records",
			ExpectedStatus:  200,
			ExpectedContent: []string{"\"items\":[]"},
			ExpectedEvents:  map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:               "org admins can list own orgs",
			Method:             http.MethodGet,
			Url:                "/api/collections/orgs/records",
			RequestHeaders:     org1AdminHeader,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		{
			Name:               "org member can list own orgs",
			Method:             http.MethodGet,
			Url:                "/api/collections/orgs/records",
			RequestHeaders:     org1MemberHeader,
			ExpectedStatus:     200,
			ExpectedContent:    []string{`"name":"org1"`, `"totalItems":1`},
			NotExpectedContent: []string{`"name":"org2"`},
			ExpectedEvents:     map[string]int{OnRecordsListRequest: 1},
			TestAppFactory:     generateTestApp,
		},
		// -----------
		// View
		// -----------
		{
			Name:            "unauth cannot view org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can view own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			RequestHeaders:  org1AdminHeader,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member can view own org",
			Method:          http.MethodGet,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			RequestHeaders:  org1MemberHeader,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"org1"`},
			ExpectedEvents:  map[string]int{OnRecordViewRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "unauth cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			RequestHeaders:  org1AdminHeader,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot create org",
			Method:          http.MethodPost,
			Url:             "/api/collections/orgs/records",
			Body:            strings.NewReader(`{"name": "newOrg"}`),
			RequestHeaders:  org1MemberHeader,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// Update
		// -----------
		{
			Name:            "unauth cannot update org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin can update own org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  org1AdminHeader,
			ExpectedStatus:  200,
			ExpectedContent: []string{`"name":"updatedOrg"`},
			ExpectedEvents:  map[string]int{OnModelAfterUpdate: 1, OnModelBeforeUpdate: 1, OnRecordAfterUpdateRequest: 1, OnRecordBeforeUpdateRequest: 1},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org members cannot update an org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  org1MemberHeader,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot update other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org2[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  org1AdminHeader,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org member cannot update other org",
			Method:          http.MethodPatch,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org2[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  org1MemberHeader,
			ExpectedStatus:  404,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		// -----------
		// Delete
		// -----------
		{
			Name:            "unauth cannot delete org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org admin cannot delete own org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			RequestHeaders:  org1AdminHeader,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
		{
			Name:            "org members cannot delete an org",
			Method:          http.MethodDelete,
			Url:             fmt.Sprintf("/api/collections/orgs/records/%s", org1[0].Id),
			Body:            strings.NewReader(`{"name": "updatedOrg"}`),
			RequestHeaders:  org1MemberHeader,
			ExpectedStatus:  403,
			ExpectedContent: []string{`"data":{}`},
			TestAppFactory:  generateTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
